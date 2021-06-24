public class InternalRepresentationUTF16 {

    private static final char[] HEX_ARRAY = "0123456789ABCDEF".toCharArray();

    public static String bytesToHex(byte[] bytes) {
        char[] hexChars = new char[bytes.length * 2];
        for (int j = 0; j < bytes.length; j++) {
            int v = bytes[j] & 0xFF;
            hexChars[j * 2] = HEX_ARRAY[v >>> 4];
            hexChars[j * 2 + 1] = HEX_ARRAY[v & 0x0F];
        }
        return new String(hexChars);
    }

    public static void main(String[] args) throws Exception {
        String s1 = new String("\u270b\u0048\u0065\u0079".getBytes("UTF-16"), "UTF-16");
        String s2 = new String("\uD83E\uDD1A\u0048\u0065\u0079".getBytes("UTF-16"), "UTF-16");
        System.out.println("✋Hey".equals(s1));
        System.out.println("🤚Hey".equals(s2));
    }
}
