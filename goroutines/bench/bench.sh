#1/bin/bash

echo ">> Benchmarks with N_THREADS = $1\n"

echo "1. C with threads"
/usr/bin/time -f "- Real %E\n- User %U\n- Sys  %S" ./exe_c $1
echo

echo "2. Go with goroutines"
/usr/bin/time -f "- Real %E\n- User %U\n- Sys  %S" ./exe_go $1
echo

echo "3. Python3 with threads"
/usr/bin/time -f "- Real %E\n- User %U\n- Sys  %S" python3 main.py $1
echo