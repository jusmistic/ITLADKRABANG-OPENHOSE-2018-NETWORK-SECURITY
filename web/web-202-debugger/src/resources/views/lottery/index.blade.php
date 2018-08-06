<!doctype html>
<html lang="{{ app()->getLocale() }}">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <title>Lottery</title>

        <!-- Fonts -->
        <link href="https://fonts.googleapis.com/css?family=Raleway:100,600" rel="stylesheet" type="text/css">

        <!-- Styles -->
        <style>
            html, body {
                background-color: #fff;
                color: #636b6f;
                font-family: 'Raleway', sans-serif;
                font-weight: 100;
                height: 100vh;
                margin: 0;
            }

            .full-height {
                height: 100vh;
            }

            .flex-center {
                align-items: center;
                display: flex;
                justify-content: center;
            }

            .position-ref {
                position: relative;
            }

            .top-right {
                position: absolute;
                right: 10px;
                top: 18px;
            }

            .content {
                text-align: center;
            }

            .title {
                font-size: 84px;
            }

            .links > a {
                color: #636b6f;
                padding: 0 25px;
                font-size: 12px;
                font-weight: 600;
                letter-spacing: .1rem;
                text-decoration: none;
                text-transform: uppercase;
            }

            .m-b-md {
                margin-bottom: 30px;
            }
        </style>
    </head>
    <body>
        <div class="flex-center position-ref full-height">
            <div class="content">
                <div class="title m-b-md">
                    ยินดีต้อนรับสู่ สลากกินไม่แบ่งเอกชน
                </div>
                @isset($is_lucky)
                @if($is_lucky)
                    <h3>Win Lucky number is {{ $number }}</h3>
                @else
                    <h3>Your number is {{ $number }} but lucky number is {{ $lucky_number }}</h3>
                @endif
                @endisset
                <div class="links">
                    <form method="POST">
                        @csrf
                        <label for="number">ชื่อของคุณ</label>
                        <input type="text" name="name" value="{{ $name or '' }}">
                        <label for="number">เลขที่ชอบ 0-10 นะจ๊ะ</label>
                        <input type="text" name="number" value="{{ $number or '' }}">
                        <button type="submit">ซื้อเลย</button>
                    </form>
                </div>
            </div>
        </div>
        <!-- Hint http://abc.com/phpmyadmin to access the database -->
    </body>
</html>
