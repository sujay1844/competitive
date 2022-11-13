#include <iostream>
#include <map>
using namespace std;

int main() {
	int t;
	cin >> t;

	map<unsigned long long int, unsigned long long int> facts;
	facts[0] = 1;
	facts[1] = 1;

	while (t--) {
		unsigned long long int n, sum=0;
		cin >> n;

		while (n--) {
			unsigned long long int x;
			cin >> x;

			if (facts.find(x) == facts.end()) {
				for (unsigned long long int i=2; i<=x ; ++i) {
					if (facts.find(i) == facts.end()) {
						facts[i] = facts[i-1] * i;
					}
				}
			}
			sum += facts[x];
		}
		sum %= 1000000007;
		cout << sum << endl;
	}
}
