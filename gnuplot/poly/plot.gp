reset
set xlabel 'time(sec)'
set ylabel 'value'
set title 'polynomial'
set term png enhanced font 'Verdana,10'
set output 'poly.png'
plot 'data.txt' using 1:2 with linespoints title 'y = x * x'

