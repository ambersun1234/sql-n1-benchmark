set title "SQL N + 1 vs. SQL JOIN benchmark testing(size 10000)"
set term png enhanced font 'Verdana,10'
set output 'benchmark.png'
set xlabel "iteration"
set ylabel "execution time(nanoseconds)"
set autoscale
set grid

plot 'n1_benchmark.txt' using 1:2 with linespoints title 'N + 1', \
'optimize_benchmark.txt' using 1:2 with linespoints title 'JOIN'