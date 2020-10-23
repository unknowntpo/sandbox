reset
set ylabel 'time(nsec)'
set style fill solid
set key right top
set title 'performance comparison - is_ascii detecting nonascii char'
set term png enhanced font 'Verdana,10'
set output 'perf_ascii.png'

plot [:][:] 'out.txt' using 1:2 with linespoint title "8 bytes/time", \
'' using 1:2 with linespoint title "1 byte/time"
