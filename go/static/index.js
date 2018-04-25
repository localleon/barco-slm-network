window.onload = function (){
    console.log('App started')
}

function apirequest(cmd, opt){
    if(opt == ''){
        fetch('/api' + '/' +  cmd , {
            method : 'GET',
            credentials : 'same-origin'
        })     
    }else{
        fetch('/api' + '/' +  cmd + '/' + opt, {
            method : 'GET',
            credentials : 'same-origin'
        })
    }
}