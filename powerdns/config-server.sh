hostnamectl set-hostname example.com

apt -y install resolvconf
systemctl stop systemd-resolved
systemctl disable systemd-resolved.service
resolvconf --enable-updates

sudo rm -f /etc/resolv.conf
echo -e "nameserver 127.0.0.1" | sudo tee /run/resolvconf/resolv.conf
sudo ln -s /run/resolvconf/resolv.conf /etc/resolv.conf
resolvconf -u
