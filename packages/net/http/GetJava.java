package chap15.network;

import static org.junit.Assert.assertEquals;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

import org.junit.Test;

public class GetJava {
	private final String USER_AGENT = "Mozilla/5.0";
	@Test
	public void testName() throws Exception {
		long start = System.nanoTime();
		
		URL obj = new URL("http://localhost:9000");
		HttpURLConnection con = (HttpURLConnection) obj.openConnection();
		con.setRequestMethod("GET");
		con.setRequestProperty("User-Agent", USER_AGENT);
		
		BufferedReader in = new BufferedReader(new InputStreamReader(con.getInputStream()));
		String inputLine;
		StringBuffer response = new StringBuffer();
		
		while ((inputLine = in.readLine()) != null) {
			response.append(inputLine);
		}
		in.close();
		System.out.println(response.toString());
		
		long end = System.nanoTime();
		System.out.println( "실행 시간 : " + ( end - start )/1000.0 );
		
	}
}
