@echo off
setlocal enabledelayedexpansion

:: Source file location
set "src_dir=gen\http"

:: Destination file location
set "dst_dir=web_page"

cd api

:: Use `goa gen` to generate service interfaces, endpoints, transport code, and OpenAPI spec.
goa gen mai.today/api/design

echo.
echo Moving OpenAPI spec files to %dst_dir%...

:: Move specific OpenAPI spec files to the destination directory
for %%f in (openapi.json openapi.yaml openapi3.json openapi3.yaml) do (
    if exist "%src_dir%\%%f" (
        move "%src_dir%\%%f" "%dst_dir%" >nul
        echo Moved %src_dir%\%%f to %dst_dir%
    ) else (
        echo File %src_dir%\%%f does not exist
    )
)

endlocal
pause
