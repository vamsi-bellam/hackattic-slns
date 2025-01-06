import java.util.Scanner;

class Solution {
  public static void main(String args[]) {
    Scanner sc = new Scanner(System.in);
    String in;
    while(sc.hasNextLine()) {
        in = sc.nextLine();
        int succCount = 1;
        String ans = "";
        for(int i = 0; i < in.length(); i++) {
            if(i+1 < in.length() && in.charAt(i) == in.charAt(i+1)) {
                succCount++;
            } else {
                if(succCount == 2) {
                    ans += String.valueOf(in.charAt(i)) + String.valueOf(in.charAt(i));
                } else if(succCount > 2) {
                    ans += succCount + String.valueOf(in.charAt(i));
                } else {
                    ans += String.valueOf(in.charAt(i));
                }
                succCount = 1;
            }
        }
        System.out.println(ans);
    }
  }
};