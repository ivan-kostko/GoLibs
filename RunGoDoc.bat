REM The repository should be in subfolder src, cause GODOC scans only src folder
godoc -http="127.0.0.1:6060" -goroot="%~dp0..\.." -notes="BUG|TODO|NOTE" -play