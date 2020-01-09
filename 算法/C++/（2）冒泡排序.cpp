#include <iostream>
#include <vector>

using namespace std;

vector<int> bubbleSort(vector<int> data) {

	int i, j, temp;

	for (i = 0; i < data.size(); i++) {

		for (j = 0; j < data.size() - i - 1; j++) {
			if (data[j] > data[j+1])
			{
				temp = data[j + 1];
				data[j + 1] = data[j];
				data[j] = temp;
			}
		}
	}
	return data;
}

int main() {

	vector<int> data;
	data = { 3,2,1 };
	data = bubbleSort(data);

	for (const int& k : data) {
		cout << k << endl;
	}

}