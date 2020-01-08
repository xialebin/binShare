#include <iostream>
#include <vector>

using namespace std;


int getKingNum(int n,int m) {
	
	//向量定义
	vector<int> arr(n);

	//初始化向量
	for (int i = 0; i < n; i++)
	{
		arr[i] = i + 1;
	}

	int startIndex = (m - 1) % n;


	while (true) {
		if (arr.size() <= 1)
		{
			break;
		}
		
		//删除操作
		arr.erase(arr.begin() + startIndex);
		startIndex = (startIndex + m - 1) % arr.size();
	}
	
	return arr[0];
}

int main() {

	int num = getKingNum(3, 2);
	cout << num << endl;
}