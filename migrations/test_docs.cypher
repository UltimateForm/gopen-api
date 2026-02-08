:begin
UNWIND [{name:"Michael Miller", properties:{debut:92.0, description:"Represent half plan film mouth former. She them education analysis cause.", id:"qxeiGXUx1GtnDfvhTD3oB7"}}, {name:"Jesse Davis", properties:{debut:51.0, description:"One relationship it difficult benefit four avoid. Operation strategy best character whole.", id:"crK6qMRExdGLrrMED97pjh"}}, {name:"Ruth Jackson", properties:{debut:9.0, description:"Car season compare whom. Skin half value high play main theory. Economy human never customer.", id:"jwzP1S4rDfEpJ2BJD4488L"}}, {name:"Deborah Cooper", properties:{debut:72.0, description:"Answer respond foot Democrat. Natural go carry not.", id:"vySzKVL9Pj2LpfvjZF36M8"}}, {name:"Brittany Preston", properties:{debut:77.0, description:"Hard whether whom audience wonder. Important billion never chance push.", id:"4QzWFid9HvaYTAwJWKgT1p"}}, {name:"Rhonda Bullock MD", properties:{debut:99.0, description:"Market able set be point deal water.", id:"rDbhox2b1yau1jx9mRzAEM"}}, {name:"Thomas Brown", properties:{debut:43.0, description:"Edge yeah land future kitchen response. Read day run office in song. Per plant travel by Mrs.", id:"c3xUXUUpibWar1emhd2oUK"}}, {name:"Randy Sullivan", properties:{debut:86.0, description:"Though serve picture. Store color back. Treat election easy knowledge sport.", id:"2qf3tqYdVyvgXs37EV26hL"}}, {name:"David Wong", properties:{debut:26.0, description:"Perform value recently war south until. Start finally send side. Born center recent structure.", id:"ojbrd5KUAsi9nAApFmXXZa"}}, {name:"Heather Ruiz", properties:{debut:99.0, description:"Single show add foot sign. Share soldier result computer number point.", id:"6BmSS5RbyYhNvucDEFZXin"}}, {name:"Cynthia Mendez", properties:{debut:79.0, description:"Window attorney nation him. Bit well assume she. Dog attorney area thousand they.", id:"fheRBjpmwp27iJSN2XXxQ9"}}, {name:"Christopher Bell", properties:{debut:53.0, description:"Community voice expert current east. Score conference it participant.", id:"jCNVTHDB88HwnTccmFAWhg"}}, {name:"Phillip Cook", properties:{debut:46.0, description:"Local result candidate much fine ground place. Bar could have maintain season memory.", id:"4zhvbmEWX5gKYq9cqwXh8D"}}, {name:"Tyler Li", properties:{debut:95.0, description:"Power performance event. Play poor pick standard soldier across quality.", id:"fYHcCFx9k7JkqQmpph4ibb"}}, {name:"Mason Martin", properties:{debut:48.0, description:"Tough employee special. So interesting rather big order threat.", id:"7a2FSFAqumC2aiqbS64YB2"}}, {name:"Kelly Green", properties:{debut:22.0, description:"Read somebody success far. Explain bill east manager decision movement once.", id:"iiV4gTSWY2X7A53A8eSeQn"}}, {name:"Jeremiah Brewer", properties:{debut:78.0, description:"Hit hit about impact political notice. If budget property stand.", id:"qxtgnL3QAZiaMTeyS1aCWz"}}, {name:"Ashley Hill", properties:{debut:7.0, description:"Performance participant any think bring discover final. Agreement activity fast order.", id:"s7gKNW9n6VaQ5fJNs4CueT"}}, {name:"Jonathan Khan", properties:{debut:25.0, description:"Here both away summer. Language national writer building physical.", id:"dWdTDRLRG4zXWkAmSgHkba"}}, {name:"Lisa Reed", properties:{debut:66.0, description:"Data example rate. Practice best me hotel.", id:"8DKbZ45ysNVwbakru1bPV1"}}] AS row
CREATE (n:Character{name: row.name}) SET n += row.properties;
UNWIND [
  {
    name: "Janice Dominguez",
    properties: {
      debut: 79.0,
      description:
        "Foreign third strategy lead listen carry plan. Down weight leader name mother rock.",
      id: "opG3HmnFgFmGZXSJDNcbFL"
    }
  },
  {
    name: "Thomas Alvarez",
    properties: {
      debut: 3.0,
      description: "Country style common. Both respond rise.",
      id: "cgi2TEUhoHANfNNkVB75Sv"
    }
  },
  {
    name: "Anita Lopez",
    properties: {
      debut: 22.0,
      description:
        "Court sometimes someone every. Term ask natural. Several language friend arm.",
      id: "obyEu7VTjFhDQ1iKCEViMk"
    }
  },
  {
    name: "Jason Brown",
    properties: {
      debut: 84.0,
      description:
        "Ten their particular boy Congress form. Attack form appear popular.",
      id: "2RU9mqXbyLsgi3KkpdRHJB"
    }
  },
  {
    name: "Brandon Blair",
    properties: {
      debut: 71.0,
      description: "Foreign almost miss weight.",
      id: "mxcrfkjwSUR3CcAKxgrdxr"
    }
  },
  {
    name: "Theodore Anderson",
    properties: {
      debut: 49.0,
      description:
        "Game still beautiful many move. Expect yard something live.",
      id: "ffnsJeLaJoe2gybZS398km"
    }
  },
  {
    name: "Mario Foley",
    properties: {
      debut: 84.0,
      description:
        "Mrs concern power industry action concern fire. Group your whom ground out clear pattern civil.",
      id: "sV3cZbw7ugudkcV1Pi88wR"
    }
  },
  {
    name: "Jennifer Hill",
    properties: {
      debut: 70.0,
      description:
        "They technology decide after. Treatment anything finish card media seem week.",
      id: "xn5mkpdvC6zrffbEhgsoGG"
    }
  },
  {
    name: "Abigail Burns",
    properties: {
      debut: 36.0,
      description:
        "Level kind parent door police road visit. Again off guy list. Little ever mean true player.",
      id: "vprNviSx5rt39X8RdkwBuE"
    }
  },
  {
    name: "Andrea Goodwin",
    properties: {
      debut: 96.0,
      description: "Bill compare best throughout fill community.",
      id: "uSxFPknefsVjJ6thpVJFZ7"
    }
  },
  {
    name: "Shawn Wu",
    properties: {
      debut: 72.0,
      description:
        "Break lay front arrive. Scene serious both strong already want. Rise million choice red include.",
      id: "11Ev9S4a7ESsMnQxHsFp2M"
    }
  },
  {
    name: "Mark Hammond",
    properties: {
      debut: 67.0,
      description: "Drive produce role current tell. Must consumer mind.",
      id: "wrDetKRTVRNUTNmT5Wnz3d"
    }
  },
  {
    name: "Timothy Harris",
    properties: {
      debut: 28.0,
      description:
        "Half involve language government environmental doctor camera responsibility.",
      id: "6KwQKXHcg8JiZGEMLidL7M"
    }
  },
  {
    name: "Sean Davis",
    properties: {
      debut: 71.0,
      description: "Trial quickly least subject.",
      id: "gc37nw1ig5C6EY6XTqRV7V"
    }
  },
  {
    name: "Christopher Dorsey",
    properties: {
      debut: 70.0,
      description:
        "Force share both protect general. Exactly development bank charge.",
      id: "gtn12jYoU8R2dKq3yaT6Tk"
    }
  },
  {
    name: "Alyssa Ross",
    properties: {
      debut: 41.0,
      description:
        "Score campaign wish pull cell meet majority. Order past window husband entire science stay.",
      id: "e3oHaoBtC6MP6LYVUc3ysp"
    }
  },
  {
    name: "Angela Spencer",
    properties: {
      debut: 25.0,
      description:
        "Smile campaign shoulder relate. Onto prevent rise itself. Beat tonight window light.",
      id: "28vmPUEWBGzgC6tF4KLQLd"
    }
  },
  {
    name: "Justin Tran",
    properties: {
      debut: 65.0,
      description:
        "Accept alone beat. Local he feel they according. Yeah environment finish hair ball task.",
      id: "2JbCBhMtd71woPt3yPaP1E"
    }
  },
  {
    name: "James Garcia",
    properties: {
      debut: 28.0,
      description: "Recently country family off world rise.",
      id: "heGX5T5L6ymLsVkTYzcQhM"
    }
  },
  {
    name: "Troy Warner",
    properties: {
      debut: 42.0,
      description:
        "Six vote crime guess. Action center nearly occur strategy relationship through.",
      id: "1gm5yCRFCursR8jv2LmpC8"
    }
  }
] AS row
CREATE (n:Character {name: row.name})
SET n += row.properties;
UNWIND [
  {
    name: "Lonnie Holt",
    properties: {
      debut: 53.0,
      description: "Help prevent range put.",
      id: "76LXxJgDdHqezdVELYkxfw"
    }
  },
  {
    name: "Nicholas Delgado",
    properties: {
      debut: 60.0,
      description: "Take design budget agree the front. Short lot hair fire.",
      id: "aBZFPLpjK3BPTg5WVwdhZR"
    }
  },
  {
    name: "Michael Whitaker",
    properties: {
      debut: 48.0,
      description:
        "Dark various already boy carry. Standard candidate interesting believe.",
      id: "tnfd9pWDvxdduT3mZUqjtB"
    }
  },
  {
    name: "Mary Russell",
    properties: {
      debut: 88.0,
      description:
        "Between thank fine down avoid option. Go into car mean available between. Fight rest fly everybody.",
      id: "shTi9VNP8sQdH49B2DLBB3"
    }
  },
  {
    name: "Leon Petersen",
    properties: {
      debut: 85.0,
      description: "On see contain fill doctor attorney.",
      id: "vZHvTkhTidfiutPExBB7Kf"
    }
  },
  {
    name: "Hannah Dorsey",
    properties: {
      debut: 67.0,
      description:
        "Test history forward enough miss machine reveal. Management realize course thousand yet while.",
      id: "fA1eMWG4geff97TR2ucJcc"
    }
  },
  {
    name: "Laura Ward",
    properties: {
      debut: 12.0,
      description: "Tv view where perhaps half you. Ready bad maintain remain.",
      id: "eZHsSMoEGHV3JTfuyst7Tb"
    }
  },
  {
    name: "Emily Stanley",
    properties: {
      debut: 63.0,
      description:
        "Base guess issue like. Score rate bill huge ago away politics.",
      id: "fFtqaQYkZJGHNTtEDfSvrR"
    }
  },
  {
    name: "Erika Vaughn",
    properties: {
      debut: 23.0,
      description:
        "Look number measure add my if expert because. Also deal avoid young. Media me realize five.",
      id: "rj4EvFTaEvba1E3BJdAF5A"
    }
  },
  {
    name: "Cristoph Rot",
    properties: {debut: 1, description: "edited", id: "d0a1f17505c56abf439835"}
  }
] AS row
CREATE (n:Character {name: row.name})
SET n += row.properties;
UNWIND [{email: "helio@mail.com", properties: {password: "q1234"}}] AS row
CREATE (n:User {email: row.email})
SET n += row.properties;:commit
:begin
UNWIND [{start: {name:"Michael Miller"}, end: {name:"Cynthia Mendez"}, properties:{level:2, spark:"Two part edge."}}, {start: {name:"Jesse Davis"}, end: {name:"Thomas Brown"}, properties:{level:2, spark:"Sure us hear."}}, {start: {name:"Jesse Davis"}, end: {name:"Brandon Blair"}, properties:{level:3, spark:"Pass success man."}}, {start: {name:"Ruth Jackson"}, end: {name:"Jeremiah Brewer"}, properties:{level:-1, spark:"Head impact follow."}}, {start: {name:"Rhonda Bullock MD"}, end: {name:"Mario Foley"}, properties:{level:2, spark:"Teach development."}}, {start: {name:"Tyler Li"}, end: {name:"Lonnie Holt"}, properties:{level:1, spark:"Leg draw Congress."}}, {start: {name:"Theodore Anderson"}, end: {name:"Andrea Goodwin"}, properties:{level:1, spark:"Assume tell too."}}, {start: {name:"Angela Spencer"}, end: {name:"James Garcia"}, properties:{level:0, spark:"Opportunity."}}, {start: {name:"Justin Tran"}, end: {name:"Mario Foley"}, properties:{level:2, spark:"Lot hand someone."}}, {start: {name:"James Garcia"}, end: {name:"Brittany Preston"}, properties:{level:1, spark:"Car pressure."}}, {start: {name:"James Garcia"}, end: {name:"Janice Dominguez"}, properties:{level:3, spark:"Beyond usually."}}, {start: {name:"Hannah Dorsey"}, end: {name:"Angela Spencer"}, properties:{level:1, spark:"Soldier strategy."}}] AS row
MATCH (start:Character{name: row.start.name})
MATCH (end:Character{name: row.end.name})
CREATE (start)-[r:FRIENDSHIP]->(end) SET r += row.properties;
UNWIND [
  {
    start: {name: "Randy Sullivan"},
    end: {name: "Andrea Goodwin"},
    properties: {level: -1, spark: "Live travel recent."}
  },
  {
    start: {name: "David Wong"},
    end: {name: "Theodore Anderson"},
    properties: {level: -1, spark: "Rather me high."}
  },
  {
    start: {name: "Heather Ruiz"},
    end: {name: "Christopher Dorsey"},
    properties: {level: 3, spark: "Sister become nor."}
  },
  {
    start: {name: "Mason Martin"},
    end: {name: "Theodore Anderson"},
    properties: {level: 2, spark: "Five size week view."}
  },
  {
    start: {name: "Kelly Green"},
    end: {name: "Mary Russell"},
    properties: {level: 2, spark: "Factor training."}
  },
  {
    start: {name: "Ashley Hill"},
    end: {name: "Brandon Blair"},
    properties: {level: -1, spark: "Outside action view."}
  },
  {
    start: {name: "Jonathan Khan"},
    end: {name: "David Wong"},
    properties: {level: -1, spark: "Remain whole effort."}
  },
  {
    start: {name: "Thomas Alvarez"},
    end: {name: "Hannah Dorsey"},
    properties: {level: -1, spark: "Opportunity quality."}
  },
  {
    start: {name: "Brandon Blair"},
    end: {name: "Jonathan Khan"},
    properties: {level: -1, spark: "System magazine."}
  },
  {
    start: {name: "Brandon Blair"},
    end: {name: "Jennifer Hill"},
    properties: {level: 3, spark: "Time throughout."}
  },
  {
    start: {name: "Brandon Blair"},
    end: {name: "Troy Warner"},
    properties: {level: 1, spark: "Country go our left."}
  },
  {
    start: {name: "Theodore Anderson"},
    end: {name: "Cynthia Mendez"},
    properties: {level: 0, spark: "Article spend child."}
  },
  {
    start: {name: "Theodore Anderson"},
    end: {name: "Jeremiah Brewer"},
    properties: {level: 3, spark: "He about series."}
  },
  {
    start: {name: "Justin Tran"},
    end: {name: "Thomas Brown"},
    properties: {level: -1, spark: "Rock realize man."}
  },
  {
    start: {name: "James Garcia"},
    end: {name: "David Wong"},
    properties: {level: -1, spark: "Manage camera seat."}
  },
  {
    start: {name: "Nicholas Delgado"},
    end: {name: "David Wong"},
    properties: {level: 0, spark: "Last east choose be."}
  },
  {
    start: {name: "Leon Petersen"},
    end: {name: "Theodore Anderson"},
    properties: {level: 3, spark: "Fine education."}
  }
] AS row
MATCH (start:Character {name: row.start.name})
MATCH (end:Character {name: row.end.name})
CREATE (start)-[r:WORSHIP]->(end)
SET r += row.properties;
UNWIND [
  {
    start: {name: "Jesse Davis"},
    end: {name: "Janice Dominguez"},
    properties: {level: 0, spark: "Crime available."}
  },
  {
    start: {name: "Heather Ruiz"},
    end: {name: "Hannah Dorsey"},
    properties: {level: 1, spark: "Individual cover."}
  },
  {
    start: {name: "Christopher Bell"},
    end: {name: "David Wong"},
    properties: {level: -1, spark: "Woman nice."}
  },
  {
    start: {name: "Christopher Bell"},
    end: {name: "Mason Martin"},
    properties: {level: 0, spark: "Region part wear."}
  },
  {
    start: {name: "Lisa Reed"},
    end: {name: "Lonnie Holt"},
    properties: {level: 1, spark: "Large watch federal."}
  },
  {
    start: {name: "Janice Dominguez"},
    end: {name: "Sean Davis"},
    properties: {level: 0, spark: "Move work record."}
  },
  {
    start: {name: "Theodore Anderson"},
    end: {name: "Heather Ruiz"},
    properties: {level: 0, spark: "Wide subject risk."}
  },
  {
    start: {name: "Abigail Burns"},
    end: {name: "Mario Foley"},
    properties: {level: 3, spark: "Consumer student."}
  },
  {
    start: {name: "Christopher Dorsey"},
    end: {name: "Rhonda Bullock MD"},
    properties: {level: 2, spark: "Hot cut write."}
  },
  {
    start: {name: "Christopher Dorsey"},
    end: {name: "Randy Sullivan"},
    properties: {level: 1, spark: "Deep buy exist."}
  },
  {
    start: {name: "James Garcia"},
    end: {name: "Mario Foley"},
    properties: {level: 3, spark: "Entire good."}
  },
  {
    start: {name: "Nicholas Delgado"},
    end: {name: "Shawn Wu"},
    properties: {level: 3, spark: "Power away how fine."}
  },
  {
    start: {name: "Mary Russell"},
    end: {name: "Brandon Blair"},
    properties: {level: 0, spark: "Up management."}
  },
  {
    start: {name: "Hannah Dorsey"},
    end: {name: "Brandon Blair"},
    properties: {level: 3, spark: "Community."}
  }
] AS row
MATCH (start:Character {name: row.start.name})
MATCH (end:Character {name: row.end.name})
CREATE (start)-[r:KILLED]->(end)
SET r += row.properties;
UNWIND [
  {
    start: {name: "Michael Miller"},
    end: {name: "Mario Foley"},
    properties: {level: 2, spark: "Although figure."}
  },
  {
    start: {name: "Jesse Davis"},
    end: {name: "Mark Hammond"},
    properties: {level: -1, spark: "Hair go several."}
  },
  {
    start: {name: "Tyler Li"},
    end: {name: "Emily Stanley"},
    properties: {level: 1, spark: "Resource eye."}
  },
  {
    start: {name: "Kelly Green"},
    end: {name: "Erika Vaughn"},
    properties: {level: -1, spark: "Executive simple."}
  },
  {
    start: {name: "Ashley Hill"},
    end: {name: "Timothy Harris"},
    properties: {level: 1, spark: "Stay down street."}
  },
  {
    start: {name: "Janice Dominguez"},
    end: {name: "Heather Ruiz"},
    properties: {level: 1, spark: "Contain attorney."}
  },
  {
    start: {name: "Thomas Alvarez"},
    end: {name: "Shawn Wu"},
    properties: {level: 3, spark: "Black morning."}
  },
  {
    start: {name: "Jason Brown"},
    end: {name: "Hannah Dorsey"},
    properties: {level: 1, spark: "Professor forget."}
  },
  {
    start: {name: "Mark Hammond"},
    end: {name: "Michael Miller"},
    properties: {level: -1, spark: "Step area low."}
  },
  {
    start: {name: "Nicholas Delgado"},
    end: {name: "Ruth Jackson"},
    properties: {level: 3, spark: "Sister fish soon."}
  },
  {
    start: {name: "Nicholas Delgado"},
    end: {name: "Jason Brown"},
    properties: {level: 1, spark: "Seem same family."}
  },
  {
    start: {name: "Michael Whitaker"},
    end: {name: "Angela Spencer"},
    properties: {level: 0, spark: "Party quickly song."}
  },
  {
    start: {name: "Mary Russell"},
    end: {name: "Jesse Davis"},
    properties: {level: 2, spark: "Miss seek."}
  },
  {
    start: {name: "Mary Russell"},
    end: {name: "Angela Spencer"},
    properties: {level: 0, spark: "Break move on cut."}
  }
] AS row
MATCH (start:Character {name: row.start.name})
MATCH (end:Character {name: row.end.name})
CREATE (start)-[r:KINSHIP]->(end)
SET r += row.properties;
UNWIND [
  {
    start: {name: "Randy Sullivan"},
    end: {name: "Jason Brown"},
    properties: {level: 3, spark: "Life allow some."}
  },
  {
    start: {name: "David Wong"},
    end: {name: "Andrea Goodwin"},
    properties: {level: 3, spark: "Certainly himself."}
  },
  {
    start: {name: "Phillip Cook"},
    end: {name: "Emily Stanley"},
    properties: {level: 0, spark: "Name world score."}
  },
  {
    start: {name: "Jeremiah Brewer"},
    end: {name: "Michael Whitaker"},
    properties: {level: 3, spark: "Fly camera eat."}
  },
  {
    start: {name: "Ashley Hill"},
    end: {name: "Justin Tran"},
    properties: {level: 0, spark: "Have my staff."}
  },
  {
    start: {name: "Jonathan Khan"},
    end: {name: "Heather Ruiz"},
    properties: {level: -1, spark: "Whatever manager."}
  },
  {
    start: {name: "Janice Dominguez"},
    end: {name: "Ruth Jackson"},
    properties: {level: 1, spark: "Sing business."}
  },
  {
    start: {name: "Mark Hammond"},
    end: {name: "Troy Warner"},
    properties: {level: 0, spark: "Focus compare this."}
  },
  {
    start: {name: "Mark Hammond"},
    end: {name: "Mary Russell"},
    properties: {level: 3, spark: "Coach study really."}
  },
  {
    start: {name: "Mark Hammond"},
    end: {name: "Laura Ward"},
    properties: {level: 2, spark: "Himself from."}
  },
  {
    start: {name: "Alyssa Ross"},
    end: {name: "Ashley Hill"},
    properties: {level: 1, spark: "Practice off me."}
  },
  {
    start: {name: "James Garcia"},
    end: {name: "Deborah Cooper"},
    properties: {level: 3, spark: "Star maintain least."}
  },
  {
    start: {name: "Troy Warner"},
    end: {name: "Randy Sullivan"},
    properties: {level: 3, spark: "Stop yeah even."}
  },
  {
    start: {name: "Lonnie Holt"},
    end: {name: "Laura Ward"},
    properties: {level: 0, spark: "Stop matter product."}
  },
  {
    start: {name: "Nicholas Delgado"},
    end: {name: "Deborah Cooper"},
    properties: {level: 0, spark: "Value check."}
  },
  {
    start: {name: "Nicholas Delgado"},
    end: {name: "Anita Lopez"},
    properties: {level: 1, spark: "And site officer."}
  },
  {
    start: {name: "Mary Russell"},
    end: {name: "Lisa Reed"},
    properties: {level: 0, spark: "Six vote cost."}
  },
  {
    start: {name: "Mary Russell"},
    end: {name: "Jennifer Hill"},
    properties: {level: 3, spark: "Cover mouth share."}
  },
  {
    start: {name: "Leon Petersen"},
    end: {name: "Michael Whitaker"},
    properties: {level: -1, spark: "They west build."}
  }
] AS row
MATCH (start:Character {name: row.start.name})
MATCH (end:Character {name: row.end.name})
CREATE (start)-[r:VASSALAGE]->(end)
SET r += row.properties;
UNWIND [
  {
    start: {name: "Ruth Jackson"},
    end: {name: "Jennifer Hill"},
    properties: {level: 3, spark: "Develop hand future."}
  },
  {
    start: {name: "Deborah Cooper"},
    end: {name: "Michael Miller"},
    properties: {level: 3, spark: "Truth response."}
  },
  {
    start: {name: "Randy Sullivan"},
    end: {name: "Ashley Hill"},
    properties: {level: 0, spark: "Cut really issue."}
  },
  {
    start: {name: "Christopher Bell"},
    end: {name: "Mary Russell"},
    properties: {level: 0, spark: "Argue he source."}
  },
  {
    start: {name: "Phillip Cook"},
    end: {name: "Justin Tran"},
    properties: {level: 2, spark: "Can cut worker line."}
  },
  {
    start: {name: "Tyler Li"},
    end: {name: "Andrea Goodwin"},
    properties: {level: 3, spark: "Memory approach."}
  },
  {
    start: {name: "Ashley Hill"},
    end: {name: "Jesse Davis"},
    properties: {level: 0, spark: "Mission fish type."}
  },
  {
    start: {name: "Lisa Reed"},
    end: {name: "Andrea Goodwin"},
    properties: {level: 0, spark: "Blood pass away one."}
  },
  {
    start: {name: "Shawn Wu"},
    end: {name: "Anita Lopez"},
    properties: {level: 2, spark: "Parent according."}
  },
  {
    start: {name: "Alyssa Ross"},
    end: {name: "Jesse Davis"},
    properties: {level: 1, spark: "Cultural list."}
  },
  {
    start: {name: "James Garcia"},
    end: {name: "Mary Russell"},
    properties: {level: -1, spark: "Economy herself."}
  },
  {
    start: {name: "Michael Whitaker"},
    end: {name: "Jeremiah Brewer"},
    properties: {level: 2, spark: "Order after parent."}
  },
  {
    start: {name: "Michael Whitaker"},
    end: {name: "Thomas Alvarez"},
    properties: {level: -1, spark: "Century without."}
  },
  {
    start: {name: "Erika Vaughn"},
    end: {name: "Troy Warner"},
    properties: {level: 3, spark: "Line beautiful."}
  }
] AS row
MATCH (start:Character {name: row.start.name})
MATCH (end:Character {name: row.end.name})
CREATE (start)-[r:ANIMOSITY]->(end)
SET r += row.properties;
UNWIND [
  {
    start: {name: "Randy Sullivan"},
    end: {name: "Laura Ward"},
    properties: {level: 2, spark: "Raise much hope."}
  },
  {
    start: {name: "Heather Ruiz"},
    end: {name: "Phillip Cook"},
    properties: {level: 0, spark: "Since fish accept."}
  },
  {
    start: {name: "Cynthia Mendez"},
    end: {name: "Tyler Li"},
    properties: {level: 0, spark: "Brother success."}
  },
  {
    start: {name: "Mark Hammond"},
    end: {name: "Lonnie Holt"},
    properties: {level: 0, spark: "Above spring cup."}
  },
  {
    start: {name: "Justin Tran"},
    end: {name: "Jennifer Hill"},
    properties: {level: 3, spark: "Walk approach bring."}
  },
  {
    start: {name: "Lonnie Holt"},
    end: {name: "Jonathan Khan"},
    properties: {level: -1, spark: "Administration meet."}
  },
  {
    start: {name: "Mary Russell"},
    end: {name: "Mason Martin"},
    properties: {level: -1, spark: "Player matter."}
  },
  {
    start: {name: "Laura Ward"},
    end: {name: "Timothy Harris"},
    properties: {level: 1, spark: "Include large only."}
  },
  {
    start: {name: "Erika Vaughn"},
    end: {name: "Christopher Dorsey"},
    properties: {level: 3, spark: "Role piece eat."}
  }
] AS row
MATCH (start:Character {name: row.start.name})
MATCH (end:Character {name: row.end.name})
CREATE (start)-[r:PEDAGOGY]->(end)
SET r += row.properties;:commit