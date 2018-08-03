var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  
  var cookie = req.cookies.user.toLowerCase();
  var visitor = '';
  
  if (cookie == 'yeri' || cookie == 'joy' || cookie == 'irene' || cookie == 'seulgi' || cookie == 'wendy' || cookie == 'cookie_jar')
  {
    visitor = cookie;
  } else {
    res.cookie("user",'visitor');
    visitor = 'visitor';
  }
  console.log(cookie);
  res.render('index', { title: visitor, jar : visitor });
});

module.exports = router;
