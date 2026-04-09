package db

import (
	"SocioGo/internal/store"
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
)

var usernames = []string{"james01","oliver99","charlotte21","daniel07","sophia12","henry34","ava56","lucas88","emma09","liam77","grace45","ethan23","amelia14","jack92","isabella33","mason11","ella27","noah44","mia08","benjamin19","harper29","samuel66","scarlett02","jacob38","hannah15","michael73","chloe25","david55","lily05","alexander91","zoe18","william42","natalie06","joseph36","madison27","christopher84","brooklyn22","matthew61","victoria13","george17","sarah26","ryan64","abigail31","nathan09","olivia47","andrew86","layla41","thomas72","bella16","charlie58","aiden35"}

var titles = []string{"10 Tips to Boost Your Productivity","The Future of Artificial Intelligence","Why Minimalism Can Change Your Life","Top 5 Programming Languages in 2025","How to Build Healthy Daily Habits","Mastering Time Management in College","The Rise of Remote Work Culture","A Beginner’s Guide to Machine Learning","How to Create a Morning Routine That Sticks","Why Reading Daily Improves Your Mindset","Exploring the Power of Mindfulness","Best Practices for Writing Clean Code","Top 10 Travel Destinations for 2025","How to Stay Motivated While Studying","The Impact of Social Media on Mental Health","Simple Recipes for Busy Professionals","Building a Personal Brand Online","Top Skills Every Developer Should Learn","How to Start Investing with Low Budget","The Secret to Building Strong Relationships","Why Consistency Beats Motivation","Understanding Cloud Computing in Simple Terms","Effective Strategies for Online Learning","How to Write Blog Posts That Rank on Google","The Benefits of Journaling Every Day","Best Free Tools for Students in 2025","How to Break Bad Habits Easily","5 Ways to Improve Your Communication Skills","The Future of Blockchain Technology","Why Sleep is the Key to Productivity","Beginner’s Guide to Freelancing","How to Overcome Procrastination","10 Simple Fitness Hacks for Beginners","Why Emotional Intelligence Matters More Than IQ","Exploring Virtual Reality in Education","How to Manage Stress Effectively","Top Mobile Apps You Need in 2025","Secrets to Writing Engaging Content","How to Develop a Growth Mindset","AI vs Human Creativity: Who Wins?","Essential Books Every Student Should Read","The Power of Networking in Your Career","How to Save Money as a College Student","Tips for Building Your First Mobile App","Why Gratitude Improves Your Happiness","The Rise of Sustainable Fashion","How to Start a Podcast from Scratch","Top Tech Trends to Watch in 2025","How to Balance Work and Personal Life"}

var contents = []string{"Discover practical strategies to maximize your daily productivity and achieve more in less time.","Explore how artificial intelligence is shaping industries and transforming everyday life.","Learn how adopting minimalism can simplify your lifestyle and boost mental clarity.","Find out which programming languages are most in demand and worth learning in 2025.","Step-by-step guide to building healthy habits that improve your body and mind.","Effective time management hacks for college students to stay ahead.","Understand how remote work is changing company culture and individual productivity.","A beginner-friendly introduction to the basics of machine learning and its applications.","Tips to design a morning routine that keeps you energized and focused.","The benefits of daily reading and how it can reshape your mindset.","Unlock the power of mindfulness for stress relief and self-growth.","Essential tips for writing clean, maintainable, and efficient code.","Discover the top travel destinations worth exploring in 2025.","Simple tricks to maintain motivation and avoid burnout while studying.","Understand the effects of social media on mental well-being.","Quick and healthy recipes for professionals with busy schedules.","Learn how to create and grow your personal brand in the digital era.","The top skills developers must master to stay competitive in tech.","How to start investing even with a small budget and grow your wealth.","The secrets to building long-lasting and meaningful relationships.","Why consistency is more important than motivation for success.","A simple explanation of cloud computing and its practical uses.","Strategies to make online learning more effective and productive.","Tips to write blog posts that rank higher in search engines.","Why journaling daily can improve your focus and emotional health.","A list of the best free tools every student should use in 2025.","Actionable tips to break bad habits and build positive ones.","Improve your communication skills with these simple techniques.","Discover the future impact of blockchain across industries.","The science of sleep and why it boosts productivity.","How to start your freelancing journey as a beginner.","Techniques to beat procrastination and stay focused.","Fitness hacks that anyone can apply to stay healthy.","Why emotional intelligence matters more than IQ in real life.","The role of virtual reality in revolutionizing education.","Effective methods to manage stress and stay balanced.","The most useful mobile apps you should download in 2025.","Proven secrets to writing engaging and compelling content.","How to develop a growth mindset and embrace challenges.","Can artificial intelligence truly replace human creativity?","Must-read books that will expand your knowledge as a student.","The importance of networking and how it helps your career.","Practical money-saving tips for college students.","Beginner-friendly steps to build your first mobile app.","Learn how gratitude can increase happiness and positivity.","How sustainable fashion is changing the clothing industry.","A complete guide to starting your own podcast.","Key technology trends that will dominate in 2025.","Tips to balance work life with personal commitments effectively."}

var tags = []string{"productivity,life hacks,success","ai,technology,future","minimalism,lifestyle,mental health","programming,languages,technology","habits,health,wellness","time management,students,success","remote work,culture,future of work","machine learning,beginner,ai","morning routine,habits,focus","reading,mindset,learning","mindfulness,mental health,focus","coding,best practices,development","travel,adventure,2025","motivation,study,students","social media,mental health,impact","recipes,food,healthy living","personal brand,career,online presence","skills,developer,career growth","investing,finance,budget","relationships,life,success","consistency,motivation,habits","cloud computing,technology,beginners","online learning,education,strategy","seo,blogging,content","journaling,mental health,habits","students,tools,free apps","habits,self improvement,change","communication,soft skills,success","blockchain,technology,future","sleep,health,productivity","freelancing,career,beginner","procrastination,focus,productivity","fitness,health,lifestyle","emotional intelligence,soft skills,career","virtual reality,education,technology","stress,wellness,mental health","mobile apps,2025,technology","content writing,blogging,marketing","growth mindset,self improvement,success","artificial"}

var comments = []string{"Great tips! I can already see these improving my productivity.","AI is really fascinating, thanks for breaking it down so clearly.","Minimalism has changed the way I think about my lifestyle.","I’m planning to learn Python next, this was helpful.","These habit tips are easy to implement and practical.","Time management advice is always useful for students like me.","Remote work is the future, this article explains it well.","Machine learning seems less intimidating after reading this.","I’m going to try creating my own morning routine now.","Reading daily is definitely improving my mindset.","Mindfulness exercises are so calming, thanks for sharing.","Clean code practices are essential for every developer.","Adding these destinations to my 2025 travel list!","Motivation tips are exactly what I needed today.","This really opened my eyes to social media’s impact.","I tried the recipes, they were super easy and tasty.","Brand building is key, great insights!","Learning these skills will definitely help my career.","Investing with a small budget is reassuring to start.","Relationship advice here is very practical and relatable.","Consistency really does make a difference, thanks!","Cloud computing seems less confusing after this.","Online learning strategies are very actionable.","I now know how to make my blog posts rank better.","Journaling has improved my focus, thanks for the tip!","These free tools are a lifesaver for students.","Breaking bad habits is tough, but these tips help.","Communication skills are so important, thanks for sharing.","Blockchain will definitely change many industries.","Sleep is so underrated, this is a good reminder.","Freelancing tips are clear and easy to follow.","Procrastination tips are going to save me so much time.","Fitness hacks are simple and effective.","Emotional intelligence is crucial, great advice!","VR in education is exciting, can’t wait to see more.","Stress management techniques are very practical.","These mobile apps are going to be super useful.","Content writing tips are very actionable and helpful.","Growth mindset advice is inspiring!","AI vs human creativity, very interesting discussion.","Book recommendations are excellent, I’ll check them out.","Networking tips are going to help my career a lot.","Money-saving tips are very practical for students.","App-building tips are simple and beginner-friendly.","Gratitude tips are very uplifting, thanks!","Sustainable fashion is something everyone should consider.","Podcast guide is very detailed and helpful.","Tech trends for 2025 are eye-opening!","Work-life balance tips are very actionable and realistic."}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx,_ := db.BeginTx(ctx, nil)
	for _, user := range users {
		if err := store.Users.Create(ctx,tx, user); err != nil {
			_= tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}
	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	cms := generateComments(500, users, posts)
	for _, comment := range cms {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding completed")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}