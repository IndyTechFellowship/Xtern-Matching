var script = //remove this top line before pasting into http://www.json-generator.com/
    [
    '{{repeat(100)}}',
    {
        _id: '{{objectId()}}',
        firstName: '{{firstName()}}',
        lastName: '{{surname()}}',
        email: '{{email()}}',
        // university: '{{company().toUpperCase()}}',
        university: function(tags) {
            var university = ['Rose-Hulman Institute of Technology', 'Indiana State University'];
            return university[tags.integer(0, university.length + 3) % university.length];
        },
        //MAJOR
        major: function(tags){
            //weight it heavier in favor of US Citizen as it is much
            var major = ['Computer Science','Computer Science', 'Computer Security', 'Software Engineering', 'Information Technology', 'Design', 'Computer Engineering'];
            return major[tags.integer(0, major.length + 3) % major.length];
        },

        gradYear:  function(tags){
            var categories = ['2017','2018'];
            return categories[tags.integer(0, categories.length - 1)];
        },
        workStatus: function(tags){
            //weight it heavier in favor of US Citizen as it is much
            var workStatus = ['US Citizen','US Citizen','Permanent Resident','H1 Visa','TN Visa','J1 Visa','F1 Visa','EAD'];
            return workStatus[tags.integer(0, workStatus.length + 2) % workStatus.length];
        },
        homeState: '{{state()}}',
        gender: '{{gender()}}',
        languages: function(tags){
            var technologies = [
                {
                    name: "Angular.js",
                    category: "Front-End"
                },
                {
                    name: "HTML",
                    category: "Front-End"
                },
                {
                    name: "JavaScript",
                    category: "Front-End"
                },	{
                    name: "CSS",
                    category: "Front-End"
                },	{
                    name: "Knockout",
                    category: "Front-End"
                },	{
                    name: "Ember",
                    category: "Front-End"
                },
                {
                    name: "Node.js",
                    category: "Full-Stack"
                },
                {
                    name: "PHP",
                    category: "Full-Stack"
                },
                {
                    name: "Ruby on Rails",
                    category: "Full-Stack"
                },
                {
                    name: ".Net",
                    category: "Full-Stack"
                },
                {
                    name: "iOS",
                    category: "Mobile"
                },
                {
                    name:"Android",
                    category:"Mobile"
                },
                {
                    name: 'C',
                    category: "General"
                },
                {
                    name:'C++',
                    category: 'General'
                },
                {
                    name:'C#',
                    category: 'General'
                },
                {
                    name:'Objective-C',
                    category: 'General'
                },
                {
                    name:'Swift',
                    category: 'General'
                },
                {
                    name:'Java',
                    category: 'General'
                },
                {
                    name:'Python',
                    category: 'General'
                },
                {
                    name:'Go',
                    category: 'General'
                },
                {
                    name:'Ruby',
                    category: 'General'
                },
                {
                    name:'SQL',
                    category: 'Database'
                },
                {
                    name:'Redis',
                    category: 'Database'
                },
                {
                    name:'Firebase',
                    category: 'Database'
                },
                {
                    name:'Mongo',
                    category: 'Database'
                },
                {
                    name:'Hadoop',
                    category: 'Database'
                },
                {
                    name:'Linux',
                    category:'Other'
                }
            ];
            return [technologies[tags.integer(0, technologies.length - 1)],
                technologies[tags.integer(0, technologies.length - 1)],
                technologies[tags.integer(0, technologies.length - 1)],
                technologies[tags.integer(0, technologies.length - 1)],
                technologies[tags.integer(0, technologies.length - 1)],
                technologies[tags.integer(0, technologies.length - 1)],
                technologies[tags.integer(0, technologies.length - 1)],
                technologies[tags.integer(0, technologies.length - 1)]];
        },
        resume: null,
        githubUrl: function(tags){
            var links = ['https://github.com/Doolan', 'https://github.com/Sp4rkfun',  'https://github.com/davisnygren', 'https://github.com/xniccum', null];
            return links[tags.integer(0, links.length -1)];
        },
        linkedinUrl: function(tags){
            var links = ['http://www.linkedin.com/', null];
            return links[tags.integer(0, links.length -1)];
        },
        personalWebiteUrl:  function(tags){
            var links = ['http://www.rose-hulman.edu/',null, null, null];
            return links[tags.integer(0, links.length - 1)];
        },
        interestedIn: function(tags){
            var categories = ['true','false'];
            return categories[tags.integer(0, categories.length - 1)];
        },
        interestedInEmail: function(tags){
            var categories = ['Designer','Software Engineer- Full Stack','Software Engineer- Back End Dev.','Software Engineer- Middle-tier Dev.','Software Engineer- Front-end Web Dev','Dev Ops','Dev Ops','Security','Product Management','Project Management'];
            return [categories[tags.integer(0, categories.length - 1)],
                categories[tags.integer(0, categories.length - 1)],
                categories[tags.integer(0, categories.length - 1)]];

        },
        //Internal Non-Suvery Attributes
        r1Grade: function (tags) {
            var grade = [
                {
                    "text":"A",
                    "value":'4'
                },{
                    "text":"B+",
                    "value":'3.5'
                },{
                    "text":"B",
                    "value":'3'
                },{
                    "text":"A",
                    "value":'4'
                },{
                    "text":"B+",
                    "value":'3.5'
                },{
                    "text":"B",
                    "value":'3'
                },{
                    "text":"B-",
                    "value":'2.8'
                },{
                    "text":"C",
                    "value":'2'
                },{
                    "text":"D",
                    "value":'1'
                }];
            return grade[tags.integer(0, grade.length - 1)];
        },
        status: function (tags) {
            var status = ['Remaining','Denied', 'Stage 1 Approved', 'Stage 1 Approved'];
            return status[tags.integer(0, status.length - 1)];
        },
        comments: [{author: '{{firstName() + " " + surname()}}', 
                    group: 'Xtern',
                    text: '{{lorem(1, "paragraphs")}}'}]
    }
];