use raylib::Color;
use raylib::consts::*;

mod animArr;

fn main() {
    let rl = raylib::init()
        .size(1600, 800)
        .title("Sort Visualiser")
        .build();
 
    rl.set_target_fps(60);

    let mut a = animArr::AnimatedArray::new(1600, 800, 2);

    while !rl.window_should_close() {
        if rl.is_key_pressed(KEY_S as i32) {
            a.shuffle(2);
		  } else if rl.is_key_pressed(KEY_ONE as i32) {
				a.bubble();
		  }

        rl.begin_drawing();
        rl.clear_background(Color::BLACK);
        a.draw(&rl);
 
        rl.draw_fps(10, 10);
        rl.end_drawing();
    }
}
