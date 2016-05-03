# The script supposed to help automaticaly generate markdown style documentation(by godocdown.exe) 
# for given folder recursively.

param
(
    [Parameter(Mandatory=$false)]
    [string]$targetPath = "" ,
    
    [Parameter(Mandatory=$false)]
    [string]$godocdown = "" 
)

# Set defaults if empty string
if ($targetPath -eq "") 
{
    $targetPath = $(Split-Path -parent $PSCommandPath)
}
if ($godocdown -eq "") 
{
    $godocdown = "$(Split-Path -parent $PSCommandPath)\godocdown.exe"
}

Function ProcessFolder($path)
{
    (get-childitem $path | where {$_.Attributes -eq 'Directory'}) | % { ProcessFolder $("$path\$_") }

    Write-Host $path
    & $godocdown "$path" > "$path\README.MD"
    
}

ProcessFolder($targetPath)
