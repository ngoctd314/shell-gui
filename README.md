# SSH GUI

## Who i am

support ssh with up/down (j/k) then enter
![sshgui demo](./result.png)

## How to download

On gateway paste (download shell gui code)

```bash
# on gateway run linux 64 bit
wget https://github.com/ngoctd314/sshgui/blob/master/shgui?raw=true && mv shgui\?raw\=true shgui && chmod +x shgui
wget https://github.com/ngoctd314/sshgui/blob/master/shgui_bar?raw=true && mv shgui_bar\?raw\=true shgui_bar && chmod +x shgui_bar
wget https://github.com/ngoctd314/sshgui/blob/master/shgui_tree?raw=true && mv shgui_tree\?raw\=true shgui_tree && chmod +x shgui_tree
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
./shgui # in case you create ssh_nav dir
./shgui -dir=custom # in case you create custom dir
```