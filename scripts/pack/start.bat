@echo off
path = %path%;.;
gin-amis-admin.exe web -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml --page ./configs/page_manager.yaml --tpl-mall ./configs/tpl_mall.yaml
pause
