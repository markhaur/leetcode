import java.util.Arrays;

public class Solution {

    public static void main(String[] args) {
        int[][] image = new int[][] {{1,1,1}, {1,1,0}, {1,0,1}};
        System.out.println("Original Image");
        for(int i =0; i < image.length; i++) {
            System.out.println(Arrays.toString(image[i]));
        }

        image = floodFill(image, 1, 1, 2);

        System.out.println("After Changing Image Color");
        for(int i = 0; i < image.length; i++) {
            System.out.println(Arrays.toString(image[i]));
        }
    }

    public static int[][] floodFill(int[][] image, int sr, int sc, int newColor) {

        if (image[sr][sc] == newColor)
            return image;
        floodFill(image, sr, sc, image[sr][sc], newColor);
        return image;
    }

    private static void floodFill(int[][] image, int sr, int sc, int color, int newColor) {
        if (sr < 0 || sc < 0 || sr >= image.length || sc >= image[0].length || image[sr][sc] != color) {
            return;
        }
        image[sr][sc] = newColor;
        floodFill(image, sr - 1, sc, color, newColor);
        floodFill(image, sr + 1, sc, color, newColor);
        floodFill(image, sr, sc - 1, color, newColor);
        floodFill(image, sr, sc + 1, color, newColor);
    }
}
