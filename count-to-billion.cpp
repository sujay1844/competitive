#include <iostream>
#include <chrono>
using namespace std;
using namespace std::chrono;

int main() {

	int order = 12;

	// Calculate the required number first
	// For some reason, looping over to create the number is faster than just typing it out.
	// Also using functions for the different steps is much slower than just doing everything in the main function.
	long long int n = 1;
	for (int i = 0; i < order ; i++) {
		n *= 10;
	}

	// Count and record the time
	auto start = high_resolution_clock::now();
	// I used int in the beginning instead long long int for i. Ran my program for a long time and nothing happened.
	// These low level languages are a pain in the ...
	long long int i = 1;
	while (i <= n){
		i++;
	}
	auto end = high_resolution_clock::now();
	auto duration  = duration_cast<microseconds>(end-start);
	float time = duration.count()/1000000.0;

	cout << "Counted to 10^" << order << endl;
	cout << "\t" << "in " << time << " sec" << endl;
	cout << "\t" << "or " << time/60.0 << " min" << endl;
	cout << "\t" << "or " << time/3600.0 << " hr" << endl;

}
/*
Output:

Counted to 10^12
	in 664.364 sec
	or 11.0727 min
	or 0.184546 hr
*/
