echo The current directory is %CD%
"C:\Program Files (x86)\Windows Kits\10\bin\10.0.18362.0\x64\SignTool.exe" sign /f .\bin\x64\Release\developer-toolkit.pfx /p "%PFX_PASSWORD%" /t http://timestamp.digicert.com scripts\install.ps1
