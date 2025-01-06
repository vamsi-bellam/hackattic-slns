
import java.lang.Math;
import java.util.Scanner;
class Solution {
  public static void main(String args[]) {
    Scanner sc = new Scanner(System.in);
    String in;
     do {
        in = sc.nextLine();
        int sum = 0;
        for(int i = 0; i< in.length(); i++) {
            if(in.charAt(i) == '#') {
                sum += (int) Math.pow(2, (in.length() - i - 1));
            }
        }
        System.out.println(sum);
    } while(sc.hasNextLine());
    
  }
};