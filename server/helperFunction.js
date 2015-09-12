/**
 *
 * Created by miloas on 2015/9/12.
 */
var fs = Npm.require('fs');
Meteor.methods({
    saveCode: function (submitedCode,pid,cpu,memory) {
        var currentUser = Meteor.user();
        var currentUserInfo = UserInfo.findOne({username:currentUser.username});
        if (currentUserInfo === undefined){
            UserInfo.insert({
                username:currentUser.username,
                solved:[]
            })
        }
        var randomCodeName = Random.id(6) + '.cc';
        fs.writeFile('D:\\'+randomCodeName,submitedCode,function(err){
            var testCasesPath = '~/testCases/'
            var cmdHead = 'ljudge';
            var cpuOption = '--max-cpu-time ' + cpu;
            var memOption = '--max-memory ' + memory + 'm';
            var codeOption = '--user-code ' + '~/submitedCode/' + randomCodeName;
            var inputOption = ' --testcase --input ' + testCasesPath+ pid + '.in';
            var outputOption = '--output ' + testCasesPath + pid + '.out';
            var cmd = [cmdHead,cpuOption,memOption,codeOption,inputOption,outputOption].join(' ');
            console.log(cmd);
       })
    }
})