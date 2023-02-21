#! /bin/bash

OutType=$1
ROOT=$PWD
nginx_log=$ROOT/bin/nginx
config_dir=$ROOT/config


if [[ $OutType == "" ]];then
    echo "useage:   ./NginxUse.sh [install,use] "
fi

if  [[ $OutType == "install" ]];then
    apt install -y nginx
    cp -r $config_dir/cert /root
    # 尝试替换日志生成位置，但是失败了
    # access_log_row=$(sed -n -e '/access_log/=' $config_dir/nginx.conf)
    # echo $access_log_row
    # access_log_message="    access_log $PWD/logs/nginx/access.log;"
    # sed -i "$access_log_row a $access_log_message" $config_dir/nginx.conf
    # sed -i "$access_log_row d" $config_dir/nginx.conf

    # error_log_row=$(sed -n -e /error_log/= $config_dir/nginx.conf)
    # echo $error_log_row
    # error_log_message="    error_log $PWD/logs/nginx/error.log;"
    # sed -i "$error_log_row a $error_log_message" $config_dir/nginx.conf
    # sed -i "$error_log_row d" $config_dir/nginx.conf
fi

if  [[ $OutType == "reload" ]];then
    cp $config_dir/default.conf /etc/nginx/conf.d
    # cp $config_dir/$config_dir/nginx.conf /etc/nginx/
    nginx -s reload
fi

if  [[ $OutType == "use" ]];then
    cp $config_dir/default.conf /etc/nginx/conf.d
    service nginx restart
    systemctl status nginx
fi

# nginx -s quit
if  [[ $OutType == "kill" ]];then
    nginx -s stop
    ps -ef | grep nginx
fi
