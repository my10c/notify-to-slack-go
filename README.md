# notify-to-slack
pipe message from nagios/naemon to a slack channel
It requires the message to be in certain format, check
how to setup your nagios/naemon notification with `notify-to-slack -s`

## usage

```
usage: notify-to-slack [-h|--help] [-v|--version] [-i|--info] [-t|--test]
                       [-s|--setup] [-S|--slack-config]

                       Simple script send a message to a slack channel via a
                       piped message.

Arguments:

  -h  --help          Print help information
  -v  --version       Show version
  -i  --info          Show how to use notify-to-slack
  -t  --test          test mode, no message will be sent. Default: false
  -s  --setup         Show how to setup in nagios or naemon
  -S  --slack-config  Show how to setup the slack configuration file
```

# how to build

```
make
```
