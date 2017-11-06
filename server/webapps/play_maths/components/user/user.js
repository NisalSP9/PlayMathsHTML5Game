/**
 * Created by nisalsp9 on 11/2/17.
 */
var user = (function () {

    var fns ={};

    fns.init = function () {

    };

    jQuery('.validatedForm').validate({
        rules: {
            "user[username]": {
                minlength: 3
            },
            "user[password]": {
                minlength: 3
            },
            "user[password_confirmation]": {
                minlength: 3,
                equalTo : "#user_password"
            }
        },
        messages: {},
        errorPlacement: function(error, element) {
            var placement = $(element).data('error');
            if (placement) {
                $(placement).append(error)
            } else {
                error.insertAfter(element);
            }
        }
    });

    $('#save').click(function () {
        if($('.validatedForm').valid()){

                alert("Ajax call");

        }else{

            alert("Fill The Form Correctly !!!");

        }
    });



return fns;

})();