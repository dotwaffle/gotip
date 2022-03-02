# gotip, but using $GOPATH

The "gotip" tool is great for using the current HEAD of the go development branch, but it does not specify where to install to. It just installs to "$HOME/sdk/gotip" without allowing any kind of configuration.

All this tool does is copy the existing [gotip](https://cs.opensource.google/go/dl/+/master:gotip/main.go;drc=294a5db0264457f6a5891594721cdd1456819de2) code but uses "$GOPATH" if it is set, which will then install there instead, which is typically "$HOME/go/sdk/gotip".

Other than this minor change, there have been no amendments made. If this tool ever breaks, I will likely re-copy the code from the latest version of [gotip](https://cs.opensource.google/go/dl/+/master:gotip/main.go;drc=294a5db0264457f6a5891594721cdd1456819de2) and make the same changes.

Good luck, have fun!
