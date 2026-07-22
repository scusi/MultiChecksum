VERSION ?= 0.9.0
APP_NAME = MultiChecksumGUI

.PHONY: build-gui-win build-gui-linux installer

# 1. GUI für Windows kompilieren (mit eingebettetem Icon)
build-gui-win:
	cd cmd/MultiChecksumGUI && \
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc \
	go build -ldflags="-H windowsgui" -o $(APP_NAME).exe main.go

# 2. Inno Setup aufrufen, um den Installer zu bauen (unter Windows oder via Wine)
#installer: build-gui-win
#	cd cmd/MultiChecksumGUI && \
#	wine iscc setup.iss

#installer: build-gui-win
#	cd cmd/MultiChecksumGUI && \
#	WINEARCH=win64 WINEPREFIX=~/.wine wine iscc setup.iss

installer: build-gui-win
	cd cmd/MultiChecksumGUI && \
	WINEARCH=win64 WINEPREFIX=~/.wine wine "C:\Program Files (x86)\Inno Setup 6\iscc.exe" setup.iss

# 3. GUI für Linux paketieren (mit Fyne CLI)
build-gui-linux:
	cd cmd/MultiChecksumGUI && \
	fyne package -os linux -name $(APP_NAME) -icon Certainity_Icon.png
