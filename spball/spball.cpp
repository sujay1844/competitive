#include <iostream>
#include <unordered_map>
using namespace std;

long fact(long n) {
	long f = 1;
	for (int i=1; i<=n ; i++) {
		f *= i;
	}
	return f;
}

int main() {
	int t;
	cin >> t;

	unordered_map<int, int> facts;

	while (t--) {
		long n, sum=0;
		cin >> n;

		while (n--) {
			long x;
			cin >> x;

			if (facts.find(x) == facts.end()) {

				for (int i=1; i<=x ; i++) {
					if (facts.find(i) == facts.end()) {
						facts[i] = fact(i);
					}
				}
			}
			sum += facts[x];

		}
		cout << sum << endl;
	}
}
