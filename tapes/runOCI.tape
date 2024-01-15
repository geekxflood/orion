Output assets/runOCI.gif

Set Theme "Dracula"
Set FontSize 15
Set Padding 20
Set Height 210

Type "docker build -t orion:latest ."
Enter
Sleep 15s

Type "docker run -d -p 9981:9981 -v config.yaml:config.yaml orion:latest run --config config.yaml"
Enter
Sleep 3s
