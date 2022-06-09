# This is the main shell file to execute Xamboo. You may change and adapt this file to your own system.
# You may also call this file directly for the systemctl and enable a daemon to control the Xamboo

# Change this to your go basic dir if needed (if you will not use default one)
# export GOPATH=/home/sites/go

# Change this to your go binary dir if needed (if you will not use default one)
# export GOBIN=/home/sites/go/bin

# Change this to your go cache dir if needed (if you will not use default one)
export GOCACHE=/home/sites/go/cache

# 1. Clean all .so if any
echo "Deletes all the old .so (they will be recompiled by Xamboo when needed)"
find . -name "*.so*" -exec rm {} \;

# 2. Run
echo "Run the xamboo"
go run xamboo.go --config=mainconfig.json --language=en
# If you compile to an executable, you may use this:
# ./xamboo --config=example/config.json
