# Get input for num
$num = Read-Host "Enter day number"

# Create folder
New-Item -Path "day_$num" -ItemType Directory 

# Create files
New-Item -Path "day_$num\input.txt" -ItemType File 
New-Item -Path "day_$num\test.txt" -ItemType File 
New-Item -Path "day_$num\problem1.md" -ItemType File 
New-Item -Path "day_$num\problem2.md" -ItemType File
New-Item -Path "day_$num\main.go" -ItemType File

Write-Host "Folder and files created successfully!"

# Pause the script execution
$key = Read-Host "Press any key to continue..."
