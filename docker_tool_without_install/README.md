-*- mode:org; fill-column:70; coding:utf-8; -*-
* Docker host
chmod 777 tool_dir

docker-compose down -v && docker-compose up -d
* Inject tools to shared folder
docker exec -it container1 sh

mkdir /home/chroot
chroot /home/chroot

ls -lth /usr/sbin/ | grep tool_dir
cp /usr/bin/python /usr/sbin/tool_dir/
* Load tool: vim
* Load tool: curl
docker exec container1 cp /usr/bin/curl /usr/sbin/tool_dir/

docker exec -it container2 sh
export PATH=$PATH:/usr/sbin/tool_dir
which curl
* load tool: python2.7
* load tool: ruby3
docker exec -it container2 sh

which python

export PATH=$PATH:/usr/sbin/tool_dir

which python
python --version
