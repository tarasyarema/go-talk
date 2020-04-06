#include "_cgo_export.h"

void lol(int n) {
    int x = Fib(n);
    // printf("fib(%2d): %d\n", n, Fib(n));
}

int fib_c(int n) {
	if (n <= 1) return 1;

	int a = 1, b = 1;

	for (int i = 2; i <= n; i++) {
		int tmp = b;
		b += a;
		a = tmp;
	}

	return b;
}