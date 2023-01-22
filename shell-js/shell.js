const inquirer = require('inquirer');
const process = require('process');
const exec = require('child_process').exec;
const readline = require('readline');
//pour controler le ctrl+p
readline.emitKeypressEvents(process.stdin);
process.stdin.setRawMode(true);
process.stdin.on('keypress', (str, key) => {
    if (key.ctrl && key.name === 'p') {
        fonctions.end.exc();
    }
});
//parseur de la command
function parser(command) {
    let args = command.split(' ');
    let cmd = args.shift();
    let ap = false;
    if (args[args.length - 1] == "!") {
        args.pop();
        ap = true;
    }
    return { cmd, args, ap };
}
//process créer sous forme de liste de processID
let processCreated = [];
//liste des différentes commands executables 
let fonctions = {
    open: {
        exc: (args, ap) => {
            if (args[0] != undefined && args.length != 1 && typeof agrs[0] == 'string') {
                exec(`open -a ${args[0]}`, () => {
                    if (!ap) {
                        exec(`ps aux | grep -v grep |grep -i ${args[0]} | awk '{print $2;}'`, (err, stdout) => {
                            let temp = stdout.split('\n')
                            temp.forEach((el) => {
                                processCreated.push(el);
                            })
                            processCreated.pop();
                        })
                    }
                })
            }
            else {
                console.error('? ' + __dirname + ' %  commande inconnue');
            }
            main();
        }
    },
    lp: {
        exc: (args, ap) => {
            if (args.length > 0) {
                exec(`ps ax`, (error, stdout) => {
                    if (error) {
                        console.error(`error: ${error.message}`);
                        return;
                    }
                    let lines = stdout.split('\n');
                    for (let i = 1; i < lines.length - 1; i++) {
                        let tmp = lines[i].split(' ');
                        tmp = tmp.filter(el => el != '');
                        console.log(`${i} nom:${tmp[4]} processId:${tmp[0]}`);
                    }
                });
            }
            else {
                console.error('? ' + __dirname + ' %  commande inconnue');
            }
            main();
        }
    },
    bing: {
        exc: (args, ap) => {
            if (args.length == 2) {
                switch (args[0]) {
                    case "-k":
                        exec(`KILL ${args[1]}`);
                        main();
                        break;
                    case "-p":
                        exec(`KILL -STOP ${args[1]}`);
                        main();
                        break;
                    case "-c":
                        exec(`KILL -CONT ${args[1]}`);
                        main();
                        break;
                    default:
                        break;
                }
            }
            else {
                console.error('? ' + __dirname + ' %  commande inconnue');
                main();
            }
        }
    },
    end: {
        exc: (args, ap) => {
            processCreated.forEach((el) => {
                exec(`KILL ${el}`);
            });
            process.exit();
        }
    },
    keep: {
        exc: (args, ap) => {
            if (args[0] != undefined && typeof args[0] == 'number' && args.length == 1) {
                exec(`nohup disown ${args[0]}`)
            }
            else {
                console.error('? ' + __dirname + ' %  commande inconnue');
            }
            main();
        }
    }

}
async function main() {
    let { command } = await inquirer.prompt([
        {
            name: 'command',
            message: __dirname + ' % '
        },
    ]);
    if (fonctions[parser(command).cmd] != undefined) {
        fonctions[parser(command).cmd].exc(parser(command).args, parser(command).ap);
    } else {
        console.error('? ' + __dirname + ' %  commande inconnue');
        main();
    }
}
main();