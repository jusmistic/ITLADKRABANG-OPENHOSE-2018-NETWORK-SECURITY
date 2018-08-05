<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\User;
use App\Lottery;

class HomeController extends Controller
{
    public function index(Request $request) {
        return view('lottery.index');
    }

    public function store(Request $request) {
        $name = $request->input('name');
        $number = $request->input('number');

        $lucky = rand(0, 10);
        if ($number == $lucky) {
            $lottery = new Lottery;
            $lottery->name = $name;
            $lottery->number = $number;
            $lottery->save();
            return view('lottery.index', ['is_lucky' => true, 'name' => $name, 'number' => $number]);
        }
        
        return view('lottery.index', ['is_lucky' => false, 'name' => $name, 'number' => $number, 'lucky_number' => $lucky]);
    }
}
