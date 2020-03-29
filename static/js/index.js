var page = {
  data: {
    url: "",
    tYear: new Date().getFullYear(),
    tMonth: new Date().getMonth() + 1,
    tDay: new Date().getDate(),
    currentMonth: new Month(new Date().getFullYear(), new Date().getMonth()),
    currentTime: new Date().getTime(),
    isLogin: false,
    token: ""
  },
  init() {
    this.setCookie("token22", "ddddddddd", 1);
    this.updateCalendar();
    this.renderCalendarTitle();
    this.bindEvent();
    this.initData();
  },
  bindEvent() {
    var _this = this;
    $(".pre").on("click", function() {
      _this.getPreMonth();
      _this.renderCalendarTitle();
      if (!_this.data.isLogin) {
        return;
      }
      var data = {
        year: _this.data.currentMonth.year,
        month: _this.data.currentMonth.month + 1,
        token: _this.data.token
      };
      _this.getListMonth(data);
    });
    $(".next").on("click", function() {
      _this.getNextMonth();
      _this.renderCalendarTitle();
      if (!_this.data.isLogin) {
        return;
      }
      var data = {
        year: _this.data.currentMonth.year,
        month: _this.data.currentMonth.month + 1,
        token: _this.data.token
      };
      _this.getListMonth(data);
    });
    $(".logout").on("click", function() {
      _this.removeCookie("token");
      _this.removeCookie("username");
      _this.data.isLogin = false;
      _this.data.token = "";
      $(".login").show();
      $(".login_ok").hide();
      $(".login_ok_username").html("");
      $(".td").removeClass("event");
      $(".list_con").html("");
    });
    $(".table").on("click", ".td", function() {
      if ($(this).hasClass("not")) {
        return;
      }
      $(this).addClass("active");
      $(".td")
        .not(this)
        .removeClass("active");
      var time = $(this).data("time");
      _this.data.currentTime = time;
      if (!_this.data.isLogin) {
        return;
      }
      var data = Object.assign({}, _this.dateFormate(), {
        token: _this.data.token
      });
      _this.getListDay(data);
    });

    $(".list_con").on("click", ".delete_btn", function() {
      var event_id = $(this).data("id");
      var data = { event_id, token: _this.data.token };
      _this.deleteEventAjax(data);
    });

    $(".login").on("click", function() {
      $(".mask").show();
      $(".login_modal").show();
      $("#login_username").val("");
      $("#login_password").val("");
      $("#login_email").val("");
    });

    $(".register_btn").on("click", function() {
      $(".mask").show();
      $(".register_modal").show();
      $(".login_modal").hide();
      $("#register_username").val("");
      $("#register_password").val("");
      $("#register_email").val("");
    });

    $(".add_btn").on("click", function() {
      if (!_this.data.isLogin) {
        alert("login");
        return;
      }
      $(".mask").show();
      $(".event_modal").show();
      $("#event_title").val("");
    });

    $(".event_modal_btn").on("click", function() {
      var title = $("#event_title").val();
      var time = $("#select").val();
      if (title === "") {
        return;
      }
      var data = Object.assign({}, _this.dateFormate(), {
        token: _this.data.token,
        title,
        time
      });
      _this.addEventAjax(data);
    });

    $(".login_modal_btn").on("click", function() {
      var username = $("#login_username").val();
      var password = $("#login_password").val();
      var email = $("#login_email").val();
      if ((username === "" || password === "", email === "")) {
        alert("Please fill in the required fields");
        return;
      }
      var data = {
        username,
        password,
        email
      };
      _this.login_ajax(data);
    });

    $(".register_modal_btn").on("click", function() {
      var username = $("#register_username").val();
      var password = $("#register_password").val();
      var email = $("#register_email").val();
      if ((username === "" || password === "", email === "")) {
        alert("Please fill in the required fields");
        return;
      }
      var data = {
        username,
        password,
        email
      };
      _this.register_ajax(data);
    });

    //close modal
    $(".modal_close_btn").on("click", function() {
      $(".modal").hide();
      $(".mask").hide();
    });
  },
  renderCalendarTitle() {
    $(".header_month").html(this.data.currentMonth.month + 1);
    $(".header_year").html(this.data.currentMonth.year);
  },
  getNextMonth() {
    this.data.currentMonth = this.data.currentMonth.nextMonth();
    this.updateCalendar();
  },
  getPreMonth() {
    this.data.currentMonth = this.data.currentMonth.prevMonth();
    this.updateCalendar();
  },
  updateCalendar() {
    var weeks = this.data.currentMonth.getWeeks();
    var html = "";
    for (var w in weeks) {
      var days = weeks[w].getDates();
      var td = "";
      for (var d in days) {
        var cYear = days[d].getFullYear();
        var cMonth = days[d].getMonth() + 1;
        var cDay = days[d].getDate();
        // console.log(days[d].getTime());
        if (
          this.data.tYear === cYear &&
          this.data.tMonth === cMonth &&
          this.data.tDay === cDay &&
          this.data.currentMonth.month + 1 === cMonth
        ) {
          td += `<td class="td active current" data-time=${days[
            d
          ].getTime()} data-day=${days[d].getDate()}>${days[d].getDate()}</td>`;
        } else if (
          this.data.tYear === cYear &&
          this.data.currentMonth.month + 1 !== cMonth
        ) {
          td += `<td class='td not' data-time=${days[
            d
          ].getTime()} data-day=${days[d].getDate()}>${days[d].getDate()}</td>`;
        } else {
          td += `<td class='td' data-time=${days[d].getTime()} data-day=${days[
            d
          ].getDate()}>${days[d].getDate()}</td>`;
        }
      }
      html += `<tr>${td}</tr>`;
    }
    $("#tbody").html(html);
  },
  initData() {
    var token = this.getCookie("token");
    var username = this.getCookie("username");
    if (!token || !username) {
      return;
    }
    this.data.isLogin = true;
    this.data.token = token;
    $(".login").hide();
    $(".login_ok").show();
    $(".login_ok_username").html(username);
    var data = Object.assign({}, this.dateFormate(), {
      token: this.data.token
    });
    this.getListMonth(data);
    this.getListDay(data);
  },
  dateFormate() {
    var time = new Date(parseInt(this.data.currentTime));
    var year = time.getFullYear();
    var month = time.getMonth() + 1;
    var day = time.getDate();
    return {
      year,
      month,
      day
    };
  },
  login_ajax(data) {
    var _this = this;
    $.ajax({
      url: this.data.url + "/user/login",
      type: "POST",
      data: data,
      success: function(res) {
        if (res.code === 0) {
          _this.setCookie("token", res.token, 1);
          _this.setCookie("username", data.username, 1);
          $(".login_modal").hide();
          $(".mask").hide();
          var username = $("#login_username").val("");
          var password = $("#login_password").val("");
          var email = $("#login_email").val("");
          _this.initData();
        } else {
          alert(res.msg);
        }
      },
      error: function() {
        alert("error");
      }
    });
  },
  register_ajax(data) {
    var _this = this;
    $.ajax({
      url: this.data.url + "/user/register",
      type: "POST",
      data: data,
      success: function(res) {
        if (res.code === 0) {
          _this.setCookie("token", res.token, 1);
          _this.setCookie("username", data.username, 1);
          $(".register_modal").hide();
          $(".mask").hide();
          var username = $("#register_username").val("");
          var password = $("#register_password").val("");
          var email = $("#register_email").val("");
          _this.initData();
        } else {
          alert(res.msg);
        }
      },
      error: function() {
        alert("error");
      }
    });
  },
  addEventAjax(data) {
    var _this = this;
    $.ajax({
      url: this.data.url + "/event/add",
      type: "POST",
      data: data,
      success: function(res) {
        if (res.code === 0) {
          $(".event_modal").hide();
          $(".mask").hide();
          _this.getListDay(data);
          _this.getListMonth(data);
        } else {
          alert(res.msg);
        }
      },
      error: function() {
        alert("error");
      }
    });
  },
  deleteEventAjax(data) {
    var _this = this;
    $.ajax({
      url: this.data.url + "/event/del",
      type: "POST",
      data: data,
      success: function(res) {
        if (res.code === 0) {
          var d = Object.assign({}, _this.dateFormate(), {
            token: _this.data.token
          });
          _this.getListDay(d);
          _this.getListMonth(d);
        } else {
          alert(res.msg);
        }
      },
      error: function() {
        alert("error");
      }
    });
  },
  getListMonth(data) {
    $.ajax({
      url: this.data.url + "/event/get_month",
      type: "POST",
      data: data,
      success: function(res) {
        if (res.code === 0) {
          var list = res.data;
          $(".td").removeClass("event");
          for (var i = 0; i < list.length; i++) {
            var els = $(".td").not(".not");
            for (var j = 0; j < els.length; j++) {
              if (els.eq(j).data("day") == list[i]) {
                els.eq(j).addClass("event");
              }
            }
          }
        } else {
          alert(res.msg);
        }
      },
      error: function() {
        alert("error");
      }
    });
  },
  getListDay(data) {
    $.ajax({
      url: this.data.url + "/event/get_day",
      type: "POST",
      data: data,
      success: function(res) {
        if (res.code === 0) {
          var list = res.data;
          var html = "";
          for (var i = 0; i < list.length; i++) {
            html += `
            <li class="event_item">
              <p class="item_title">${list[i].title}</p>
              <span class="item_time">${list[i].time}</span>
              <div class="delete_btn" data-id=${list[i].event_id}>x</div>
          </li>
            `;
          }
          $(".list_con").html(html);
        } else {
          alert(res.msg);
        }
      },
      error: function() {
        alert("error");
      }
    });
  },
  setCookie(name, value, iDay) {
    var oDate = new Date();
    oDate.setDate(oDate.getDate() + iDay);
    document.cookie = name + "=" + value + ";expires=" + oDate;
  },
  /*****获取cookie*****/
  getCookie(name) {
    var arr = document.cookie.split("; ");
    for (var i = 0; i < arr.length; i++) {
      var arr2 = arr[i].split("=");
      if (arr2[0] == name) {
        return arr2[1];
      }
    }
    return "";
  },
  /*****移除cookie*****/
  removeCookie(name) {
    this.setCookie(name, 1, -1);
  }
};
$(function() {
  page.init();
});
