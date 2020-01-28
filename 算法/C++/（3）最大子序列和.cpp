#include <iostream>
#include <vector>

using namespace std;

int getMaxNum(vector<int> data,int left,int right) {

	if (left == right)
	{
		return data[left];
	}

	int midden;
	midden = (left + right) / 2;

	int leftMax = 0;
	int sum = 0;
	for (int i = midden;i >= left; i--)
	{
		sum += data[i];
		if (sum > leftMax)
		{
			leftMax = sum;
		}
	}

	int rightMax = 0;
	sum = 0;
	for (int i = midden+1; i <= right; i++)
	{
		sum += data[i];
		if (sum > rightMax)
		{
			rightMax = sum;
		}
	}

	int leftNum,rightNum;
	leftNum = getMaxNum(data, left, midden);
	rightNum = getMaxNum(data, midden + 1, right);

	int max;
	max = leftNum > rightNum ? leftNum : rightNum;
	if (max > leftMax+rightMax)
	{
		return max;
	}
	else
	{
		return leftMax + rightMax;
	}
}

int main() {

	vector<int> data = {1,2,-3};
	int num = getMaxNum(data, 0, data.size() - 1);
	cout << num << endl;
}