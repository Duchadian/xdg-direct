#!/bin/bash
sudo cp ./config/xdg-direct.desktop /usr/share/applications/xdg-direct.desktop
xdg-settings set default-url-scheme-handler http xdg-direct.desktop
xdg-settings set default-url-scheme-handler https xdg-direct.desktop
