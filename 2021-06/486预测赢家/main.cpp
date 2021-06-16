//我的方法：暴力法，只要有一种情况 玩家1 肯定胜利即可
class Solution {
public:
    bool PredictTheWinner(vector<int>& nums) {
        int n = nums.size();
        return helper(nums,0,n-1,0,0);
    }

    bool helper(vector<int>& nums,int left,int right,int score1,int score2){
        if(left > right) {
            return score1 >= score2;
        }
        if (left == right){
            return score1+nums[left] >= score2;
        }
        

        return (helper(nums,left+1,right-1,score1+nums[left],score2+nums[right]) && 
                helper(nums,left+2,right,score1+nums[left],score2+nums[left+1])     )  ||  
               (helper(nums,left+1,right-1,score1+nums[right],score2+nums[left]) && 
                helper(nums,left,right-2,score1+nums[right],score2+nums[right-1]) );

    }
};

//动态规划 dp
//之前的递归存在重复子问题    如第一回合取了left、right 第二回合取了left+1 right-1
//                      等价于第一回合取了left、left+1，第二回合取了right、right-1，
//                          最后都只剩下 [left+2,right-2]
class Solution {
public:
    bool PredictTheWinner(vector<int>& nums) {
        //dp方法
        int n = nums.size();
        vector<vector<int>> dp(n,vector<int>(n,0));
        //边界条件
        for(int i = 0;i<n;i++){
            dp[i][i] = nums[i];
        }

        for(int i = n-2;i>=0;i--){
            for(int j = i+1;j<n;j++){
                dp[i][j] = max(nums[j] - dp[i][j-1],nums[i]-dp[i+1][j]);
            }
        }
        return dp[0][n-1] >= 0;
    }
};