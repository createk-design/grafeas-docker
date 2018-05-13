package main

import (
	"context"
	//"log"

	pb "github.com/grafeas/grafeas/v1alpha1/proto"
	"google.golang.org/grpc"
	"github.com/grafeas/client-go/v1alpha1"
	"fmt"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGrafeasClient(conn)

	my_note_request := pb.CreateNoteRequest{}
	my_note := pb.Note{}
	my_note.Name = "projects/project_one/notes/note_one"
	my_note_request.Note = &my_note

	_, err = client.CreateNote(context.Background(), &my_note_request)

	if err != nil {
		log.Fatalf("Error creating note %v", err)
	} else {
		log.Printf("Succesfully created note")
	}

	//nPID := "best-vuln-scanner"
	//nID := "CVE-2014-9911"
	//n := note(nPID, nID)
	//createdN, err := client.CreateNote(context.Background(), nPID, nID, *n)
	//if err != nil {
	//	log.Fatalf("Error creating note %v", err)
	//} else {
	//	log.Printf("Succesfully created note: %v", createdN)
	//}

	// List notes
	//resp, err := client.ListNotes(context.Background(),
	//	&pb.ListNotesRequest{
	//		Parent: "projects/myproject",
	//	})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//if len(resp.Notes) != 0 {
	//	log.Println(resp.Notes)
	//} else {
	//	log.Println("Project does not contain any notes")
	//}
}


func note(pID, nID string) *v1alpha1.Note {
	return &v1alpha1.Note{
		Name:             fmt.Sprintf("projects/%v/notes/%v", pID, nID),
		ShortDescription: "CVE-2014-9911",
		LongDescription:  "NIST vectors: AV:N/AC:L/Au:N/C:P/I:P",
		Kind:             "PACKAGE_VULNERABILITY",
		VulnerabilityType: v1alpha1.VulnerabilityType{
			CvssScore: 7.5,
			Severity:  "HIGH",
			Details: []v1alpha1.Detail{
				{
					CpeUri:   "cpe:/o:debian:debian_linux:7",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: v1alpha1.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: v1alpha1.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:7",
						Package_: "icu",
						Version: v1alpha1.Version{
							Name:     "4.8.1.1",
							Revision: "12+deb7u6",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:debian:debian_linux:8",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: v1alpha1.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: v1alpha1.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: v1alpha1.Version{
							Name:     "52.1",
							Revision: "8+deb8u4",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:debian:debian_linux:9",
					Package_: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: v1alpha1.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: v1alpha1.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:9",
						Package_: "icu",
						Version: v1alpha1.Version{
							Name:     "55.1",
							Revision: "3",
						},
					},
				},
				{
					CpeUri:   "cpe:/o:canonical:ubuntu_linux:14.04",
					Package_: "andriod",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: v1alpha1.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "MEDIUM",

					FixedLocation: v1alpha1.VulnerabilityLocation{
						CpeUri:   "cpe:/o:canonical:ubuntu_linux:14.04",
						Package_: "andriod",
						Version: v1alpha1.Version{
							Kind: "MAXIMUM",
						},
					},
				},
			},
		},
		RelatedUrl: []v1alpha1.RelatedUrl{
			{
				Url:   "https://security-tracker.debian.org/tracker/CVE-2014-9911",
				Label: "More Info",
			},
			{
				Url:   "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9911",
				Label: "More Info",
			},
		},
	}
}

func Occurrence(noteName string, occurrenceName string) *v1alpha1.Occurrence {
	return &v1alpha1.Occurrence{
		Name: occurrenceName,
		ResourceUrl: "gcr.io/foo/bar",
		NoteName:    noteName,
		Kind:        "PACKAGE_VULNERABILITY",
		VulnerabilityDetails: v1alpha1.VulnerabilityDetails{
			Severity:  "HIGH",
			CvssScore: 7.5,
			PackageIssue: []v1alpha1.PackageIssue{
				v1alpha1.PackageIssue{
					SeverityName: "HIGH",
					AffectedLocation: v1alpha1.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: v1alpha1.Version{
							Name:     "52.1",
							Revision: "8+deb8u3",
						},
					},
					FixedLocation: v1alpha1.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: v1alpha1.Version{
							Name:     "52.1",
							Revision: "8+deb8u4",
						},
					},
				},
			},
		},
	}
}
