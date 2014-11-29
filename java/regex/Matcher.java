import java.util.regex.Pattern;

public class Matcher {

    private static final Pattern SIMPLE_URL_PATTERN =
	Pattern.compile("(\\w+://)(.+@)*([\\w\\d\\.]+)(:[\\d]+){0,1}/*(.*)");

    public static boolean isValid(String url) {
	return SIMPLE_URL_PATTERN.matcher(url).matches();
    }

    public static void main(String[] args) {
	if (isValid(args[0])) {
	    System.out.println("PASS: matched");
	} else {
	    System.out.println("FAIL: not mached");
	}
    }
}
