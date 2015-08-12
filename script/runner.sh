# python cef_parser.py -r -v --go_file=cef.go --c_file=cef.c main.h -I../ > dump.js
python cef_parser.py --go_file=../cef.go --c_file=../cef.c --h_file=cef.h --h_path=../ main.h -I../ 
tput bel