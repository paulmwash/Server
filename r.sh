CONN=RADIENT 001
sudo nmcli connection modify "" ipv4.dns "1.1.1.1 1.0.0.1" ipv4.ignore-auto-dns yes
sudo nmcli connection up ""
sudo systemctl restart nscd dnsmasq systemd-resolved 2>/dev/null || true
