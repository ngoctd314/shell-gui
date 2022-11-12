# SSH GUI

## Who i am

support ssh with up/down (j/k) then enter
![sshgui demo](./result.png)

## How to download

On gateway paste (download sshgui code)

```bash
# on gateway run linux 64 bit
wget https://github.com/ngoctd314/sshgui/blob/master/run_linux?raw=true && mv run\?raw\=true sshgui && chmod +x sshgui

# on gateway run macos 64 bit
wget https://github.com/ngoctd314/sshgui/blob/master/run_mac?raw=true && mv run\?raw\=true sshgui && chmod +x sshgui

# on gateway run win 64 bit
wget https://github.com/ngoctd314/sshgui/blob/master/run_win?raw=true && mv run\?raw\=true sshgui && chmod +x sshgui
```


Create ssh folder navigation
```bash
cd
mkdir ssh_nav # or any name you want
```

Then create your own ssh server
```txt
├── ssh_nav
│   ├── ssp
│   │   └── ssp@192.168.1.1_2395
│   └── trino
│       └── trino@192.168.1.1_2395
```

Run
```bash
cd
./sshgui # in case you create ssh_nav dir
./sshgui -dir=custom # in case you create custom dir
```