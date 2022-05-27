apt -y install resolvconf
systemctl stop systemd-resolved
systemctl disable systemd-resolved.service

sudo rm -f /etc/resolv.conf
echo -e "nameserver 1.1.1.1\nsearch ." | sudo tee /run/resolvconf/resolv.conf
sudo ln -s /run/resolvconf/resolv.conf /etc/resolv.conf

hostnamectl set-hostname example.com
echo -e "nameserver 1.1.1.1" | sudo tee /etc/resolvconf/resolv.conf.d/head > /dev/null
cat /etc/resolvconf/resolv.conf.d/head
resolvconf --enable-updates
resolvconf -u
