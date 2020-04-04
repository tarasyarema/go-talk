#include <bits/stdc++.h>
#include <thread>
#include <atomic>

using namespace std;

#define MAX_T 90000
#define DEF_T 1000
#define LOOPS 1000

atomic<int> counter(0);

void example()
{
    counter++;
}

int main(int argc, char *argv[])
{
    int n = (argc > 1) ? atoi(argv[1]) : DEF_T;

    vector<thread> ts;

    for (int i = 0; i < n; i++)
        ts.push_back(thread(example));

    for (auto &t : ts)
        t.join();

    cout << "> the counter is " << counter << endl;

    return 0;
}