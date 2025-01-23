use std::net::UdpSocket;
//use std::thread;
//use std::time;

fn main() {
    let socket: UdpSocket = UdpSocket::bind( "0.0.0.0:0").expect("Couldnt bind");
    //lag melding: 

    let buf:[u8; 3] = [3,2,3];
    
    socket.send_to(&buf, "10.24.36.154:20000").expect("Couldnt send");
    //thread::sleep(time::Duration::from_millis(10));
    //socket.send_to(&buf, "10.24.36.154:20000").expect("Couldnt send");
 
    
}
