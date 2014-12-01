package nak3.com;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;

/**
 * Servlet implementation class dbTest
 */
@WebServlet("/testServlet")
public class ServletLog4j extends HttpServlet {
    private static final long serialVersionUID = 1L;

    /**
     * @see HttpServlet#HttpServlet()
     */
    public ServletLog4j() {
	super();
	// TODO Auto-generated constructor stub
    }

    /**
     * @see HttpServlet#doGet(HttpServletRequest request, HttpServletResponse
     *      response)
     */
    protected void doGet(HttpServletRequest request,
			 HttpServletResponse response) throws ServletException, IOException {
	System.out.println("request forwarding ...");

	for (int i = 0; i < 100; i++) {
	    Logger logger = Logger.getLogger("operation");
	    logger.fatal("test operation fatal");
	    logger.error("test operation error");
	    logger.warn("test operation warn");
	    logger.info("test operation info");
	    logger.debug("test operation debug");
	    logger.trace("test operation trace");

	}
	response.sendRedirect("forwardingTest.html");
    }
}
