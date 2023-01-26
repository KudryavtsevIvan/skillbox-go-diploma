# Service provider system
Итоговая работа по курсу Skillbox "Go-разработчик"


## Сборка
Используя go

    go build cmd/main.go



## Сторонние Пакеты
    github.com/gorilla/mux - мультиплексор HTTP-запросов
    github.com/go-yaml/yaml - поддержка YAML для языка Go
    



## Запуск

Для корректной работы нужно запустить симулятор и отдельно запустить основную программу

Запуск симулятора

    go run simulator/main.go

Запуск основной программы 

    go run cmd/main.go




## Запрос
Пример запроса на сервер

    http://localhost:8282/

Пример ответа от сервера


{
   "sms": [
      [
         {
            "country": " Austria",
            "bandwidth": "71",
            "response_time": "164",
            "provider": "Topolo"
         },
         {
            "country": " Bulgaria",
            "bandwidth": "80",
            "response_time": "1398",
            "provider": "Rond"
         },
         {
            "country": " Canada",
            "bandwidth": "56",
            "response_time": "339",
            "provider": "Rond"
         },
         {
            "country": " Denmark",
            "bandwidth": "41",
            "response_time": "1432",
            "provider": "Topolo"
         },
         {
            "country": " France",
            "bandwidth": "29",
            "response_time": "163",
            "provider": "Topolo"
         },
         {
            "country": " Monaco",
            "bandwidth": "17",
            "response_time": "1805",
            "provider": "Kildy"
         },
         {
            "country": " New Zealand",
            "bandwidth": "62",
            "response_time": "1140",
            "provider": "Kildy"
         },
         {
            "country": " Peru",
            "bandwidth": "73",
            "response_time": "1191",
            "provider": "Topolo"
         },
         {
            "country": " Russia",
            "bandwidth": "89",
            "response_time": "549",
            "provider": "Topolo"
         },
         {
            "country": " Spain",
            "bandwidth": "72",
            "response_time": "1247",
            "provider": "Topolo"
         },
         {
            "country": " Switzerland",
            "bandwidth": "31",
            "response_time": "1893",
            "provider": "Topolo"
         },
         {
            "country": " Turkey",
            "bandwidth": "25",
            "response_time": "246",
            "provider": "Rond"
         },
         {
            "country": " United States",
            "bandwidth": "49",
            "response_time": "49",
            "provider": "Rond"
         }
      ],
      [
         {
            "country": " Austria",
            "bandwidth": "71",
            "response_time": "164",
            "provider": "Topolo"
         },
         {
            "country": " Bulgaria",
            "bandwidth": "80",
            "response_time": "1398",
            "provider": "Rond"
         },
         {
            "country": " Canada",
            "bandwidth": "56",
            "response_time": "339",
            "provider": "Rond"
         },
         {
            "country": " Denmark",
            "bandwidth": "41",
            "response_time": "1432",
            "provider": "Topolo"
         },
         {
            "country": " France",
            "bandwidth": "29",
            "response_time": "163",
            "provider": "Topolo"
         },
         {
            "country": " Monaco",
            "bandwidth": "17",
            "response_time": "1805",
            "provider": "Kildy"
         },
         {
            "country": " New Zealand",
            "bandwidth": "62",
            "response_time": "1140",
            "provider": "Kildy"
         },
         {
            "country": " Peru",
            "bandwidth": "73",
            "response_time": "1191",
            "provider": "Topolo"
         },
         {
            "country": " Russia",
            "bandwidth": "89",
            "response_time": "549",
            "provider": "Topolo"
         },
         {
            "country": " Spain",
            "bandwidth": "72",
            "response_time": "1247",
            "provider": "Topolo"
         },
         {
            "country": " Switzerland",
            "bandwidth": "31",
            "response_time": "1893",
            "provider": "Topolo"
         },
         {
            "country": " Turkey",
            "bandwidth": "25",
            "response_time": "246",
            "provider": "Rond"
         },
         {
            "country": " United States",
            "bandwidth": "49",
            "response_time": "49",
            "provider": "Rond"
         }
      ]
   ],
   "mms": [
      [
         {
            "country": " Monaco",
            "provider": "Kildy",
            "bandwidth": "80",
            "response_time": "56"
         },
         {
            "country": " Monaco",
            "provider": "Kildy",
            "bandwidth": "80",
            "response_time": "56"
         },
         {
            "country": " New Zealand",
            "provider": "Kildy",
            "bandwidth": "80",
            "response_time": "1205"
         },
         {
            "country": " Bulgaria",
            "provider": "Rond",
            "bandwidth": "49",
            "response_time": "1495"
         },
         {
            "country": " Canada",
            "provider": "Rond",
            "bandwidth": "57",
            "response_time": "1174"
         },
         {
            "country": " Turkey",
            "provider": "Rond",
            "bandwidth": "83",
            "response_time": "736"
         },
         {
            "country": " United States",
            "provider": "Rond",
            "bandwidth": "10",
            "response_time": "1518"
         },
         {
            "country": " Austria",
            "provider": "Topolo",
            "bandwidth": "89",
            "response_time": "831"
         },
         {
            "country": " Denmark",
            "provider": "Topolo",
            "bandwidth": "88",
            "response_time": "129"
         },
         {
            "country": " France",
            "provider": "Topolo",
            "bandwidth": "70",
            "response_time": "930"
         },
         {
            "country": " Great Britain",
            "provider": "Topolo",
            "bandwidth": "54",
            "response_time": "489"
         },
         {
            "country": " Peru",
            "provider": "Topolo",
            "bandwidth": "38",
            "response_time": "268"
         },
         {
            "country": " Russia",
            "provider": "Topolo",
            "bandwidth": "29",
            "response_time": "860"
         },
         {
            "country": " Spain",
            "provider": "Topolo",
            "bandwidth": "0",
            "response_time": "682"
         },
         {
            "country": " Switzerland",
            "provider": "Topolo",
            "bandwidth": "92",
            "response_time": "497"
         }
      ],
      [
         {
            "country": " Monaco",
            "provider": "Kildy",
            "bandwidth": "80",
            "response_time": "56"
         },
         {
            "country": " Monaco",
            "provider": "Kildy",
            "bandwidth": "80",
            "response_time": "56"
         },
         {
            "country": " New Zealand",
            "provider": "Kildy",
            "bandwidth": "80",
            "response_time": "1205"
         },
         {
            "country": " Bulgaria",
            "provider": "Rond",
            "bandwidth": "49",
            "response_time": "1495"
         },
         {
            "country": " Canada",
            "provider": "Rond",
            "bandwidth": "57",
            "response_time": "1174"
         },
         {
            "country": " Turkey",
            "provider": "Rond",
            "bandwidth": "83",
            "response_time": "736"
         },
         {
            "country": " United States",
            "provider": "Rond",
            "bandwidth": "10",
            "response_time": "1518"
         },
         {
            "country": " Austria",
            "provider": "Topolo",
            "bandwidth": "89",
            "response_time": "831"
         },
         {
            "country": " Denmark",
            "provider": "Topolo",
            "bandwidth": "88",
            "response_time": "129"
         },
         {
            "country": " France",
            "provider": "Topolo",
            "bandwidth": "70",
            "response_time": "930"
         },
         {
            "country": " Great Britain",
            "provider": "Topolo",
            "bandwidth": "54",
            "response_time": "489"
         },
         {
            "country": " Peru",
            "provider": "Topolo",
            "bandwidth": "38",
            "response_time": "268"
         },
         {
            "country": " Russia",
            "provider": "Topolo",
            "bandwidth": "29",
            "response_time": "860"
         },
         {
            "country": " Spain",
            "provider": "Topolo",
            "bandwidth": "0",
            "response_time": "682"
         },
         {
            "country": " Switzerland",
            "provider": "Topolo",
            "bandwidth": "92",
            "response_time": "497"
         }
      ]
   ],
   "voice_call": [
      {
         "country": "RU",
         "bandwidth": "30",
         "response_time": "745",
         "provider": "TransparentCalls",
         "connection_stability": 0.7,
         "ttfb": 484,
         "voice_purity": 25,
         "median_of_calls_time": 57
      },
      {
         "country": "GB",
         "bandwidth": "93",
         "response_time": "372",
         "provider": "TransparentCalls",
         "connection_stability": 0.83,
         "ttfb": 676,
         "voice_purity": 79,
         "median_of_calls_time": 50
      },
      {
         "country": "FR",
         "bandwidth": "18",
         "response_time": "436",
         "provider": "TransparentCalls",
         "connection_stability": 0.71,
         "ttfb": 207,
         "voice_purity": 53,
         "median_of_calls_time": 55
      },
      {
         "country": "AT",
         "bandwidth": "11",
         "response_time": "657",
         "provider": "TransparentCalls",
         "connection_stability": 0.66,
         "ttfb": 578,
         "voice_purity": 15,
         "median_of_calls_time": 39
      },
      {
         "country": "BG",
         "bandwidth": "16",
         "response_time": "531",
         "provider": "E-Voice",
         "connection_stability": 0.83,
         "ttfb": 455,
         "voice_purity": 86,
         "median_of_calls_time": 37
      },
      {
         "country": "DK",
         "bandwidth": "88",
         "response_time": "137",
         "provider": "JustPhone",
         "connection_stability": 0.9,
         "ttfb": 613,
         "voice_purity": 9,
         "median_of_calls_time": 54
      },
      {
         "country": "CA",
         "bandwidth": "85",
         "response_time": "157",
         "provider": "JustPhone",
         "connection_stability": 0.67,
         "ttfb": 360,
         "voice_purity": 20,
         "median_of_calls_time": 46
      },
      {
         "country": "ES",
         "bandwidth": "65",
         "response_time": "1776",
         "provider": "E-Voice",
         "connection_stability": 0.67,
         "ttfb": 264,
         "voice_purity": 61,
         "median_of_calls_time": 6
      },
      {
         "country": "CH",
         "bandwidth": "28",
         "response_time": "1989",
         "provider": "JustPhone",
         "connection_stability": 0.71,
         "ttfb": 786,
         "voice_purity": 44,
         "median_of_calls_time": 35
      },
      {
         "country": "TR",
         "bandwidth": "92",
         "response_time": "1725",
         "provider": "TransparentCalls",
         "connection_stability": 0.94,
         "ttfb": 178,
         "voice_purity": 28,
         "median_of_calls_time": 9
      },
      {
         "country": "PE",
         "bandwidth": "87",
         "response_time": "78",
         "provider": "JustPhone",
         "connection_stability": 0.83,
         "ttfb": 32,
         "voice_purity": 84,
         "median_of_calls_time": 56
      },
      {
         "country": "NZ",
         "bandwidth": "68",
         "response_time": "1838",
         "provider": "JustPhone",
         "connection_stability": 0.99,
         "ttfb": 695,
         "voice_purity": 21,
         "median_of_calls_time": 51
      },
      {
         "country": "MC",
         "bandwidth": "7",
         "response_time": "159",
         "provider": "E-Voice",
         "connection_stability": 0.67,
         "ttfb": 819,
         "voice_purity": 34,
         "median_of_calls_time": 48
      }
   ],
   "email": {
      "AT": [
         [
            {
               "country": "AT",
               "provider": "Gmail",
               "delivery_time": 35
            },
            {
               "country": "AT",
               "provider": "Comcast",
               "delivery_time": 51
            },
            {
               "country": "AT",
               "provider": "Live",
               "delivery_time": 59
            }
         ],
         [
            {
               "country": "AT",
               "provider": "Hotmail",
               "delivery_time": 506
            },
            {
               "country": "AT",
               "provider": "GMX",
               "delivery_time": 569
            },
            {
               "country": "AT",
               "provider": "Yandex",
               "delivery_time": 594
            }
         ]
      ],
      "BG": [
         [
            {
               "country": "BG",
               "provider": "Hotmail",
               "delivery_time": 27
            },
            {
               "country": "BG",
               "provider": "Protonmail",
               "delivery_time": 85
            },
            {
               "country": "BG",
               "provider": "Mail.ru",
               "delivery_time": 124
            }
         ],
         [
            {
               "country": "BG",
               "provider": "Yahoo",
               "delivery_time": 388
            },
            {
               "country": "BG",
               "provider": "Live",
               "delivery_time": 512
            },
            {
               "country": "BG",
               "provider": "Yandex",
               "delivery_time": 555
            }
         ]
      ],
      "CA": [
         [
            {
               "country": "CA",
               "provider": "MSN",
               "delivery_time": 71
            },
            {
               "country": "CA",
               "provider": "GMX",
               "delivery_time": 115
            },
            {
               "country": "CA",
               "provider": "Yahoo",
               "delivery_time": 161
            }
         ],
         [
            {
               "country": "CA",
               "provider": "Protonmail",
               "delivery_time": 396
            },
            {
               "country": "CA",
               "provider": "Live",
               "delivery_time": 539
            },
            {
               "country": "CA",
               "provider": "Hotmail",
               "delivery_time": 577
            }
         ]
      ],
      "CH": [
         [
            {
               "country": "CH",
               "provider": "MSN",
               "delivery_time": 62
            },
            {
               "country": "CH",
               "provider": "Live",
               "delivery_time": 194
            },
            {
               "country": "CH",
               "provider": "GMX",
               "delivery_time": 271
            }
         ],
         [
            {
               "country": "CH",
               "provider": "RediffMail",
               "delivery_time": 449
            },
            {
               "country": "CH",
               "provider": "Yahoo",
               "delivery_time": 454
            },
            {
               "country": "CH",
               "provider": "Mail.ru",
               "delivery_time": 462
            }
         ]
      ],
      "DK": [
         [
            {
               "country": "DK",
               "provider": "MSN",
               "delivery_time": 44
            },
            {
               "country": "DK",
               "provider": "AOL",
               "delivery_time": 44
            },
            {
               "country": "DK",
               "provider": "Gmail",
               "delivery_time": 69
            }
         ],
         [
            {
               "country": "DK",
               "provider": "GMX",
               "delivery_time": 479
            },
            {
               "country": "DK",
               "provider": "Mail.ru",
               "delivery_time": 529
            },
            {
               "country": "DK",
               "provider": "Hotmail",
               "delivery_time": 588
            }
         ]
      ],
      "ES": [
         [
            {
               "country": "ES",
               "provider": "Orange",
               "delivery_time": 13
            },
            {
               "country": "ES",
               "provider": "RediffMail",
               "delivery_time": 138
            },
            {
               "country": "ES",
               "provider": "Yandex",
               "delivery_time": 178
            }
         ],
         [
            {
               "country": "ES",
               "provider": "Protonmail",
               "delivery_time": 432
            },
            {
               "country": "ES",
               "provider": "Mail.ru",
               "delivery_time": 434
            },
            {
               "country": "ES",
               "provider": "Yahoo",
               "delivery_time": 597
            }
         ]
      ],
      "FR": [
         [
            {
               "country": "FR",
               "provider": "Orange",
               "delivery_time": 220
            },
            {
               "country": "FR",
               "provider": "MSN",
               "delivery_time": 233
            },
            {
               "country": "FR",
               "provider": "GMX",
               "delivery_time": 233
            }
         ],
         [
            {
               "country": "FR",
               "provider": "Yandex",
               "delivery_time": 421
            },
            {
               "country": "FR",
               "provider": "Gmail",
               "delivery_time": 518
            },
            {
               "country": "FR",
               "provider": "Comcast",
               "delivery_time": 585
            }
         ]
      ],
      "GB": [
         [
            {
               "country": "GB",
               "provider": "Mail.ru",
               "delivery_time": 25
            },
            {
               "country": "GB",
               "provider": "Live",
               "delivery_time": 64
            },
            {
               "country": "GB",
               "provider": "GMX",
               "delivery_time": 110
            }
         ],
         [
            {
               "country": "GB",
               "provider": "Hotmail",
               "delivery_time": 345
            },
            {
               "country": "GB",
               "provider": "Yandex",
               "delivery_time": 515
            },
            {
               "country": "GB",
               "provider": "MSN",
               "delivery_time": 572
            }
         ]
      ],
      "MC": [
         [
            {
               "country": "MC",
               "provider": "Gmail",
               "delivery_time": 38
            },
            {
               "country": "MC",
               "provider": "AOL",
               "delivery_time": 60
            },
            {
               "country": "MC",
               "provider": "MSN",
               "delivery_time": 62
            }
         ],
         [
            {
               "country": "MC",
               "provider": "Mail.ru",
               "delivery_time": 529
            },
            {
               "country": "MC",
               "provider": "Comcast",
               "delivery_time": 553
            },
            {
               "country": "MC",
               "provider": "Hotmail",
               "delivery_time": 570
            }
         ]
      ],
      "NZ": [
         [
            {
               "country": "NZ",
               "provider": "RediffMail",
               "delivery_time": 36
            },
            {
               "country": "NZ",
               "provider": "Mail.ru",
               "delivery_time": 79
            },
            {
               "country": "NZ",
               "provider": "Protonmail",
               "delivery_time": 97
            }
         ],
         [
            {
               "country": "NZ",
               "provider": "Gmail",
               "delivery_time": 461
            },
            {
               "country": "NZ",
               "provider": "Yandex",
               "delivery_time": 525
            },
            {
               "country": "NZ",
               "provider": "Comcast",
               "delivery_time": 590
            }
         ]
      ],
      "PE": [
         [
            {
               "country": "PE",
               "provider": "Gmail",
               "delivery_time": 9
            },
            {
               "country": "PE",
               "provider": "MSN",
               "delivery_time": 172
            },
            {
               "country": "PE",
               "provider": "Comcast",
               "delivery_time": 184
            }
         ],
         [
            {
               "country": "PE",
               "provider": "GMX",
               "delivery_time": 499
            },
            {
               "country": "PE",
               "provider": "Orange",
               "delivery_time": 586
            },
            {
               "country": "PE",
               "provider": "Mail.ru",
               "delivery_time": 596
            }
         ]
      ],
      "RU": [
         [
            {
               "country": "RU",
               "provider": "Protonmail",
               "delivery_time": 85
            },
            {
               "country": "RU",
               "provider": "RediffMail",
               "delivery_time": 101
            },
            {
               "country": "RU",
               "provider": "Mail.ru",
               "delivery_time": 138
            }
         ],
         [
            {
               "country": "RU",
               "provider": "Gmail",
               "delivery_time": 464
            },
            {
               "country": "RU",
               "provider": "Comcast",
               "delivery_time": 523
            },
            {
               "country": "RU",
               "provider": "AOL",
               "delivery_time": 583
            }
         ]
      ],
      "TR": [
         [
            {
               "country": "TR",
               "provider": "Hotmail",
               "delivery_time": 90
            },
            {
               "country": "TR",
               "provider": "Protonmail",
               "delivery_time": 93
            },
            {
               "country": "TR",
               "provider": "Live",
               "delivery_time": 169
            }
         ],
         [
            {
               "country": "TR",
               "provider": "AOL",
               "delivery_time": 561
            },
            {
               "country": "TR",
               "provider": "Yahoo",
               "delivery_time": 564
            },
            {
               "country": "TR",
               "provider": "RediffMail",
               "delivery_time": 569
            }
         ]
      ],
      "US": [
         [
            {
               "country": "US",
               "provider": "Gmail",
               "delivery_time": 143
            },
            {
               "country": "US",
               "provider": "RediffMail",
               "delivery_time": 200
            },
            {
               "country": "US",
               "provider": "GMX",
               "delivery_time": 206
            }
         ],
         [
            {
               "country": "US",
               "provider": "AOL",
               "delivery_time": 484
            },
            {
               "country": "US",
               "provider": "MSN",
               "delivery_time": 541
            },
            {
               "country": "US",
               "provider": "Yahoo",
               "delivery_time": 574
            }
         ]
      ]
   },
   "billing": {
      "create_customer": false,
      "purchase": true,
      "payout": false,
      "recurring": true,
      "fraud_control": true,
      "checkout_page": false
   },
   "support": [
      3,
      84
   ],
   "incident": [
      {
         "topic": "MMS connection stability",
         "status": "closed"
      },
      {
         "topic": "Checkout page is down",
         "status": "closed"
      },
      {
         "topic": "Buy phone number not working in US",
         "status": "closed"
      },
      {
         "topic": "API Slow latency",
         "status": "closed"
      },
      {
         "topic": "API Slow latency",
         "status": "closed"
      },
      {
         "topic": "API Slow latency",
         "status": "closed"
      }
   ]
}
