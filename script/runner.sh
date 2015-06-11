# python auto_bind.py ../include/capi/cef_app_capi.h -I../> dump.txt
# python cef_parser.py main.h -I../ > dump.js
# python cef_parser.py main.h -I../ > dump.txt
python cef_parser.py main.h -I../
tput bel