use raylib::Color;
use raylib::Vector2;

use std::ops::{Index, IndexMut};

extern crate rand;
use rand::Rng;

pub struct AnimatedArray {
    data: Vec<f32>,
    line_width: i32,
    line_num: i32,
    max_value: f32,
    width: i32,
    height: i32,
    accesses: u32,
    active: usize,
    active2: usize,
    sorted: bool,
    shuffling: bool
}

impl Index<usize> for AnimatedArray {
    type Output = f32;

    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.data[i]
    }
}

impl IndexMut<usize> for AnimatedArray {    // Do not use when just drawing
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        self.accesses += 1;
        println!("Accesses: {}", self.accesses);
        &mut self.data[i]
    }
}

impl AnimatedArray {
    pub fn new(w: i32, h: i32, line_w: i32) -> AnimatedArray {
        let mut a = AnimatedArray {
            data: vec![],
            line_width: line_w,
            line_num: w/line_w,
            max_value: -1.0,
            width: w,
            height: h,
            accesses: 0,
            active: 0,
            active2: 0,
            sorted: true,
            shuffling: false,
        };

        a.generate_data(0.0, a.height as f32, (a.height/a.line_num) as f32);
        return a;
    }

    fn reset_values(&mut self) {
        self.accesses = 0;
        self.active = 0;
        self.active2 = 0;
        self.sorted = false;
        self.shuffling = false;
    }

    pub fn draw(&self, rl: &raylib::RaylibHandle) {
        for i in (0..self.data.len()) {
            let mut col = Color::RED;
            if i == self.active {
            	col = Color::GREEN;
				} else if i == self.active2 {
					col = Color::RED;
				}
            self.draw_line(rl, i, col);
        }
    }

    fn draw_line(&self, rl: &raylib::RaylibHandle, i: usize, col: Color) {
        let x: f32 = (i as f32 * self.line_width as f32) + (self.line_width as f32/2.0);
        let y: f32 = self.height as f32 - self.data[i];
        rl.draw_line_ex(Vector2 { x: x, y: self.height as f32 }, Vector2 { x: x, y: y }, self.line_width as f32, col);
    }

    fn generate_data(&mut self, start: f32, finish: f32, jump: f32) {
        if self.data.len() > 0 {
            self.data = vec![];
        }

        self.max_value = finish;
        let mut i = start + jump;
        while i <= finish {
            self.data.push(i);
            i += jump
        }
    }

    fn swap_elements(&mut self, in1: usize, in2: usize) {
        self.data.swap(in1, in2);
        self.accesses += 2
    }

    pub fn shuffle(&mut self, passes: u8) {
        self.sorted = false;
        self.shuffling = true;
        
        for _ in (0..passes) {
            for j in (0..self.data.len()) {
                self.active = j;
                self.active2 = rand::thread_rng().gen_range(0, self.data.len());
                self.swap_elements(self.active, self.active2)
            }
        }
        self.reset_values();
    }

	 pub fn bubble(&mut self) {
		self.sorted = false;
		while !self.sorted {
			self.sorted = true;
			for i in (0..self.data.len()-1) {
				self.active = i;
				self.active2 = i + 1;
				if self.data[self.active] > self.data[self.active2] {
					self.swap_elements(self.active, self.active2);
					self.sorted = false;
				}
			}
		}

		self.reset_values();
	 }
}
