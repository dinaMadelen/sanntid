use std::net::UdpSocket; // https://doc.rust-lang.org/std/net/struct.UdpSocket.html
use std::thread;


fn main(){

    let socket = UdpSocket::bind("10.22.107.12:20071").expect("couldn't bind to address");
    socket.connect("10.22.107.12:20071");
    println!("Connected to server at 10.22.107.12:20071");

    let target_address = " 10.22.107.12:20070";

    let socket_clone = socket.try_clone().expect("Failed to clone the socket");

    let messages = vec!["Test"];

    //Recive
    thread::spawn(move || {
        socket.send_to(&buf)
        
    });


    //Sending
    loop {
        thread::sleep(std::time::Duration::from_secs(1));
    }
    

}