#!/bin/sh

set -u

main() {
    # Exit if `cobi` is already installed
    if check_cmd cobi; then
        err "cobi already installed on this machine"
    fi

    # Start installing
    echo "Installing cobi ..."

    # Check system information
    ostype="$(uname -s | tr '[:upper:]' '[:lower:]')"
    cputype="$(uname -m | tr '[:upper:]' '[:lower:]')"
    check_architecture "$ostype" "$cputype"
    progressBar 10 100

    # Initialization
    ensure mkdir -p "$HOME/.cobi/bin"

    # Install terraform
    if [ $cputype = "x86_64" ];then
      cputype="amd64"
    fi
    if [ $cputype = "aarch64" ];then
      cputype="arm"
    fi
    progressBar 30 100

    # Download cobi binary
    cobi_url="https://cobi-releases.s3.ap-south-1.amazonaws.com/cobi_${ostype}_${cputype}"
    ensure downloader "$cobi_url" "$HOME/.cobi/bin/cobi"
    ensure chmod +x "$HOME/.cobi/bin/cobi"
    progressBar 90 100

    # Try adding the cobi directory to PATH
    add_path
    progressBar 100 100
    sleep 1

    # Output success message
    printf "\n\n"
    printf 'If you are using a custom shell, make sure you update your PATH.\n'
    printf "     export PATH=\$PATH:\$HOME/.cobi/bin"
    printf "\n\n"
    printf "Done! Restart terminal and run the command below to begin.\n"
    printf "\n"
    printf "cobi --help\n"
}

# Check prerequisites for installing cobi-cli.
prerequisites() {
    # Check commands
    need_cmd uname
    need_cmd chmod
    need_cmd mkdir
    need_cmd rm

    # Install unzip for user if not installed
    if ! check_cmd unzip; then
        echo "installing prerequisites: unzip"
        if ! sudo apt-get install unzip -qq; then
             err "need 'unzip' (command not found)"
        fi
    fi

    # Check either curl or wget is installed.
    if ! check_cmd curl; then
        if ! check_cmd wget; then
          err "need 'curl' or 'wget' (command not found)"
        fi
    fi
}

# Check if cobi-cli supports given system and architecture.
check_architecture() {
    ostype="$1"
    cputype="$2"

    if [ "$ostype" = 'linux' -a "$cputype" = 'x86_64' ]; then
        :
    elif [ "$ostype" = 'linux' -a "$cputype" = 'aarch64' ]; then
        :
    elif [ "$ostype" = 'darwin' ]; then
        if [ "$cputype" = 'x86_64' ]; then
            :
        elif [ "$cputype" = 'arm64' ]; then
            :
        else
            echo 'unsupported OS type or architecture'
            exit 1
        fi

        case $(sw_vers -productVersion) in
            10.*)
                # If we're running on macOS, older than 10.13, then we always
                # fail to find these options to force fallback
                if [ "$(sw_vers -productVersion | cut -d. -f2)" -lt 13 ]; then
                    # Older than 10.13
                    echo "Warning: Detected macOS platform older than 10.13"
                    return 1
                fi
                ;;
            11.*)
                # We assume Big Sur will be OK for now
                ;;
            12.*)
                # We assume Monterey will be OK for now
                ;;
            13.*)
                # We assume Ventura will be OK for now
                ;;
            14.*)
                # We assume Sonoma will be OK for now
                ;;
            *)
                # Unknown product version, warn and continue
                echo "Warning: Detected unknown macOS major version: $(sw_vers -productVersion)"
                echo "Warning TLS capabilities detection may fail"
                ;;
       esac
    else
        echo 'unsupported OS type or architecture'
        exit 1
    fi
}

# Add the binary path to $PATH.
add_path(){
    if ! check_cmd cobi; then
        if [ -f "$HOME/.zprofile" ] ; then
            echo "" >> "$HOME/.zprofile"
            echo 'export PATH=$PATH:$HOME/.cobi/bin' >> "$HOME/.zprofile"
        else
            echo 'export PATH=$PATH:$HOME/.cobi/bin' > "$HOME/.zprofile"
        fi

        if [ -f "$HOME/.bash_profile" ] ; then
            echo "" >> "$HOME/.bash_profile"
            echo 'export PATH=$PATH:$HOME/.cobi/bin' >> "$HOME/.bash_profile"
        else
            echo 'export PATH=$PATH:$HOME/.cobi/bin' > "$HOME/.bash_profile"
        fi

        if [ -f "$HOME/.cshrc" ] ; then
            echo "" >> "$HOME/.cshrc"
            echo 'setenv PATH $PATH\:$HOME/.cobi/bin' >> "$HOME/.cshrc"
        fi

        echo "" >> "$HOME/.profile"
        echo 'export PATH=$PATH:$HOME/.cobi/bin' >> "$HOME/.profile"
    fi
}

# Source: https://sh.rustup.rs
check_cmd() {
    command -v "$1" > /dev/null 2>&1
}

# Source: https://sh.rustup.rs
need_cmd() {
    if ! check_cmd "$1"; then
        err "need '$1' (command not found)"
    fi
}

# Source: https://sh.rustup.rs
err() {
    echo ''
    echo "$1" >&2
    exit 1
}

# Source: https://sh.rustup.rs
ensure() {
    if ! "$@"; then err "command failed: $*"; fi
}

# This wraps curl or wget. Try curl first, if not installed, use wget instead.
# Source: https://sh.rustup.rs
downloader() {
    if check_cmd curl; then
        if ! check_help_for curl --proto --tlsv1.2; then
            curl --silent --show-error --fail --location "$1" --output "$2"
        else
            curl --proto '=https' --tlsv1.2 --silent --show-error --fail --location "$1" --output "$2"
        fi
    elif check_cmd wget; then
        if ! check_help_for wget --https-only --secure-protocol; then
            wget "$1" -O "$2"
        else
            wget --https-only --secure-protocol=TLSv1_2 "$1" -O "$2"
        fi
    else
        echo "Unknown downloader"   # should not reach here
    fi
}

# Source: https://sh.rustup.rs
check_help_for() {
    local _cmd
    local _arg
    local _ok
    _cmd="$1"
    _ok="y"
    shift

    for _arg in "$@"; do
        if ! "$_cmd" --help | grep -q -- "$_arg"; then
            _ok="n"
        fi
    done

    test "$_ok" = "y"
}

# Source: https://github.com/fearside/ProgressBar
progressBar() {
    _progress=$1
    _done=$((_progress*5/10))
    _left=$((50-_done))
    done=""
    if ! [ $_done = "0" ];then
        done=$(printf '#%.0s' $(seq $_done))
    fi
    left=""
    if ! [ $_left = "0" ];then
      left=$(printf '=%.0s' $(seq $_left))
    fi
    printf "\rProgress : [$done$left] ${_progress}%%"
}

main "$@" || exit 1
