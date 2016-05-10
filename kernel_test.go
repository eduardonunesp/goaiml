package goaiml

import "testing"

func Test_Kernel_Loader_From_File(t *testing.T) {
	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromFile("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}
}

func Test_Kernel_Learn_Tag(t *testing.T) {
	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromFile("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	v, err := aiml.Respond("load aiml b")
	if err != nil {
		t.Error(err)
	}

	if v != "Ready now :D" {
		t.Error("Result not match")
	}

	v, err = aiml.Respond("just load")
	if err != nil {
		t.Error(err)
	}

	if v != "YEAH MAN" {
		t.Error("Result not match")
	}
}

func Test_Kernel_Loader(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
		    <category>
		        <pattern>
		        	HI
		        </pattern>
		        <template>
		        	HELLO!
		        </template>
		    </category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}
}

func Test_Kernel_Respond(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
		    <category>
		        <pattern>
		        	HI
		        </pattern>
		        <template>
		        	HELLO!
		        </template>
		    </category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("HI")
	if err != nil {
		t.Error(err)
	}

	if result != "HELLO!" {
		t.Log(result)
		t.Error("Result not match")
	}
}

func Test_Kernel_Respond_Star(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
				<pattern>
					MY DOGS NAME IS *
				</pattern>
				<template>
				    That is interesting that you have a dog named <star />
				</template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("My DOGs Name is Bela")
	if err != nil {
		t.Error(err)
	}

	if result != "That is interesting that you have a dog named Bela" {
		t.Error("Result not match:", result)
	}
}

func Test_Kernel_Respond_Star_Star(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
				<pattern>
					MY DOGS NAME IS * AND *
				</pattern>
				<template>
				    That is interesting that you have a dog named <star /> and <star />
				</template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("My DOGs Name is Bela and Bruce")
	if err != nil {
		t.Error(err)
	}

	if result != "That is interesting that you have a dog named Bela and Bruce" {
		t.Error("Result not match:", result)
	}
}

func Test_Kernel_Respond_Memory(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
				<pattern>MY DOGS NAME IS *</pattern>
				<template>
				    That is interesting that you have a dog named
				    <set name="dog">
				        <star />
				    </set>
				</template>
			</category>
			<category>
			    <pattern>
			    	WHAT IS MY DOGS NAME
			    </pattern>
			    <template>
			        Your dog's name is <get name="dog" />
			    </template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("My DOGs Name is Bela")
	if err != nil {
		t.Error(err)
	}

	if result != "That is interesting that you have a dog named Bela" {
		t.Log(result)
		t.Error("Result not match")
	}

	result, err = aiml.Respond("WHaT IS My DOGS nAME")
	if err != nil {
		t.Error(err)
	}

	if result != "Your dog's name is Bela" {
		t.Log(result)
		t.Error("Result not match")
	}
}

func Test_Kernel_Respond_Bot_At_Template(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
			    <pattern>
			    	DO YOU HAVE ANY IDEA
			    </pattern>
			    <template>
			        No, I'm <bot name="name" />
			    </template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("Do you have any idea")
	if err != nil {
		t.Error(err)
	}

	if result != "No, I'm "+BOT_NAME {
		t.Log(result)
		t.Error("Result not match")
	}
}

func Test_Kernel_Respond_Bot_At_Pattern(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
			    <pattern>
			    	<bot name="name"/>
			    	*
			    </pattern>
			    <template>
			        No, I'm <bot name="name" />
			    </template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond(BOT_NAME + " are you there")
	if err != nil {
		t.Error(err)
	}

	if result != "No, I'm "+BOT_NAME {
		t.Log(result)
		t.Error("Result not match")
	}
}

func Test_Kernel_Respond_At_Think(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
			    <pattern>I AM *</pattern>
			    <template>
			        <think>
			        	<set name="ok">
			        		<star />
			        	</set>
			        </think>
			        Maybe :D
			    </template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("I AM YOUR FRIEND")
	if err != nil {
		t.Error(err)
	}

	if result != "Maybe :D" {
		t.Log(result)
		t.Error("Result not match")
	}
}

func Test_Kernel_Respond_At_Srai(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
			    <pattern>I AM *</pattern>
			    <template>
			        <think>
			        	<set name="ok">
			        		<star />
			        	</set>
			        </think>
			        Maybe :D
			    </template>
			</category>
			<category>
			    <pattern>
			    	DO I KNOW
			    </pattern>
			    <template>
					<srai>I AM IRON MAN</srai>
			    </template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("DO I KNOW")
	if err != nil {
		t.Error(err)
	}

	if result != "Maybe :D" {
		t.Error("Result not match")
	}
}

func Test_Kernel_Respond_At_Random(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
			    <pattern>DO YOU THINK</pattern>
			    <template>
					<random>
						<li>
							Why are you Joking
						</li>
						<li>
							Do your friends call you
						</li>
						<li>
							My name is ` + BOT_NAME + `
						</li>
					</random>
			    </template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("Do you think")
	if err != nil {
		t.Error(err)
	}

	if result != "Why are you Joking" && result != "Do your friends call you" && result != "My name is GOAIMLBot" {
		t.Error("Result not match:", result)
	}
}

func Test_Kernel_Respond_At_Condition(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
			    <pattern>DO YOU THINK</pattern>
			    <template>
					<condition>
						<li name="key" value="value">
							Why are you Joking
						</li>
						<li>
							Do your friends call you
						</li>
					</condition>
			    </template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	aiml.Memory["key"] = "value"
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("Do you think")
	if err != nil {
		t.Error(err)
	}

	if result != "Why are you Joking" {
		t.Error("Result not match:", result)
	}
}

func Test_Kernel_Respond_At_Input(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<category>
				<pattern>YES</pattern>
				<template>Understood</template>
			</category>
			<category>
				<pattern>SURE</pattern>
				<template>Yes, you said <input index="2" /></template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("YES")
	if err != nil {
		t.Error(err)
	}

	result, err = aiml.Respond("SURE")
	if err != nil {
		t.Error(err)
	}

	if result != "Yes, you said YES" {
		t.Error("Result not match:", result)
	}
}

func Test_Kernel_Respond_At_Topic_1(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<topic name="dogs">
				<category>
					<pattern>DO YOU LIKE DOGS</pattern>
					<template>I LIKE</template>
				</category>
			</topic>
			<category>
				<pattern>DO YOU LIKE DOGS</pattern>
				<template>Nope</template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	aiml.Memory["topic"] = "dogs"
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("Do you like dogs")
	if err != nil {
		t.Error(err)
	}

	if result != "I LIKE" {
		t.Error("Result not match:", result)
	}
}

func Test_Kernel_Respond_At_Topic_2(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<topic name="dogs">
				<category>
					<pattern>DO YOU LIKE DOGS</pattern>
					<template>I LIKE</template>
				</category>
			</topic>
			<category>
				<pattern>* DOGS *</pattern>
				<template>
					<think>
						<set name="topic">dogs</set>Ok, Dogs
					</think>
				</template>
			</category>
			<category>
				<pattern>DO YOU LIKE DOGS</pattern>
				<template>Nope</template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("Let's talk about dogs")
	if err != nil {
		t.Error(err)
	}

	result, err = aiml.Respond("Do you like dogs")

	if result != "I LIKE" {
		t.Error("Result not match:", result)
	}
}

func Test_Kernel_Respond_At_Topic_3(t *testing.T) {
	xml := []byte(
		`<aiml version="1.0.1" encoding="UTF-8">
			<topic name="dogs">
				<category>
					<pattern>* DOGS *</pattern>
					<template>
						<set name="topic"></set>
						Ok
					</template>
				</category>
			</topic>
			<category>
				<pattern>DO YOU LIKE DOGS</pattern>
				<template>Nope</template>
			</category>
		</aiml>`,
	)

	aiml := NewAIMLInterpreter()
	aiml.Memory["topic"] = "dogs"
	err := aiml.LearnFromXML(xml)
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("Let's talk about dogs")
	if err != nil {
		t.Error(err)
	}

	result, err = aiml.Respond("Do you like dogs")

	if result != "Nope" {
		t.Error("Result not match:", result)
	}
}
